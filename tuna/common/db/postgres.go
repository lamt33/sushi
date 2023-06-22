package db

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lamt3/sushi/tuna/common/logger"
	_ "github.com/lib/pq"
)

func PGBuilder() *postgresBuilder {
	return &postgresBuilder{}
}

type postgresBuilder struct {
	connectionString string
	fn               func(db *sqlx.DB)
}

func (pc *postgresBuilder) ConnectionString(connectionString string) *postgresBuilder {
	pc.connectionString = connectionString
	return pc
}

func (pc *postgresBuilder) Connection(pgHost, dbPort, pgUser, pgPW, pgDB string) *postgresBuilder {
	pc.connectionString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgHost,
		dbPort,
		pgUser,
		pgPW,
		pgDB,
	)
	return pc
}

func (pc *postgresBuilder) AddSettings(fn func(db *sqlx.DB)) *postgresBuilder {
	pc.fn = fn
	return pc
}

func (pc *postgresBuilder) Build() *sqlx.DB {

	db, err := sqlx.Open("postgres", pc.connectionString)
	if err != nil {
		log.Fatalf("could not open sql connection : %v", err)
	}
	for i := 0; i < 4; i++ {
		err = db.Ping()
		if err == nil {
			break
		} else {
			if i == 3 {
				log.Fatalf("could not open sql connection : %v", err)
			}
			log.Printf("could not ping DB %v. retrying", err)
			time.Sleep(time.Second * 10)
			continue
		}
	}

	pc.fn(db)
	log.Println("DB Is Connected...")
	return db
}

// Use if need primary and secondary set up
func OPPGBuilder() *opPostgresBuilder {
	return &opPostgresBuilder{}
}

type opPostgresBuilder struct {
	primaryConn      string
	primarySetting   func(db *sqlx.DB)
	secondaryConn    []string
	secondarySetting func(db *sqlx.DB)
}

func (pc *opPostgresBuilder) PrimaryConn(connectionString string, addPrimarySettingsFn func(db *sqlx.DB)) *opPostgresBuilder {
	pc.primaryConn = connectionString
	pc.primarySetting = addPrimarySettingsFn
	return pc
}

func (pc *opPostgresBuilder) SecondaryConn(connectionString []string, addSecondarySettingsFn func(db *sqlx.DB)) *opPostgresBuilder {
	pc.secondaryConn = connectionString
	pc.secondarySetting = addSecondarySettingsFn
	return pc
}

func (pc *opPostgresBuilder) Build() *OPPG {

	//Build Primary
	if pc.primaryConn == "" {
		log.Fatalf("no primary conn")
	}
	primaryDB := pc.build(pc.primaryConn, pc.primarySetting)

	//Build Secondary
	secondaryDBs := []*sqlx.DB{}
	if pc.secondaryConn != nil && len(pc.secondaryConn) > 0 {
		for _, s := range pc.secondaryConn {
			secondaryDBs = append(secondaryDBs, pc.build(s, pc.secondarySetting))
		}
	}

	return &OPPG{
		primary:   primaryDB,
		secondary: secondaryDBs,
		count:     1,
	}
}

func (pc *opPostgresBuilder) build(conn string, fn func(db *sqlx.DB)) *sqlx.DB {
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		log.Fatalf("could not open sql connection : %v", err)
	}

	for i := 0; i < 4; i++ {
		err = db.Ping()
		if err == nil {
			break
		} else {
			if i == 3 {
				log.Fatalf("could not open sql connection : %v", err)
			}
			log.Printf("could not ping DB %v. retrying", err)
			time.Sleep(time.Second * 10)
			continue
		}
	}

	fn(db)
	logger.Info("DB Is Connected...")
	return db
}

type OPPG struct {
	primary   *sqlx.DB
	secondary []*sqlx.DB
	count     uint64 // Monotonically incrementing counter on each query
}

func (op *OPPG) Primary() *sqlx.DB {
	return op.primary
}

func (op *OPPG) QueryRow(query string, args ...interface{}) *sqlx.Row {
	return op.Secondary().QueryRowx(query, args)
}

func (op *OPPG) Secondary() *sqlx.DB {
	return op.secondary[op.findSecondary(len(op.secondary))]
}

func (op *OPPG) findSecondary(n int) int {
	if op.count == 9223372036854775807 { //reset count to prevent int overflow
		op.count = 1
	}

	return int((atomic.AddUint64(&op.count, 1) % uint64(n-1)))
}

func (op *OPPG) Close() error {
	err := op.primary.Close()
	if err != nil {
		return logger.Error("could not close primary db %+v", err)
	}

	for _, s := range op.secondary {
		err := s.Close()
		if err != nil {
			return logger.Error("could not close secondary db %+v", err)
		}
	}
	return nil
}

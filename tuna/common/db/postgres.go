package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
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

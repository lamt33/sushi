package repo

import (
	"context"

	"github.com/lamt3/sushi/tuna/common/db"
)

type IAccountRepo interface {
	InsertUser()
	UpdateUser()
}

type acccountRepo struct {
	DB *db.OPPG
}

func NewPGAccountRepo(pgDB *db.OPPG) *acccountRepo {
	return &acccountRepo{
		DB: pgDB,
	}

}

func (pr *acccountRepo) InsertUser() {
	primary := pr.DB.Primary()
	primary.ExecContext(context.Background(), "INSERT")
}

func (pr *acccountRepo) UpdateUser() {

}

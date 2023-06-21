package repo

import (
	"github.com/jmoiron/sqlx"
)

type IAccountRepo interface {
	InsertUser()
	UpdateUser()
}

type acccountRepo struct {
	DB *sqlx.DB
}

func NewPGAccountRepo(pgDB *sqlx.DB) *acccountRepo {
	return &acccountRepo{
		DB: pgDB,
	}

}

func (pr *acccountRepo) InsertUser() {

}

func (pr *acccountRepo) UpdateUser() {

}

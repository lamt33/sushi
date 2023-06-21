package account

import (
	"github.com/lamt3/sushi/tuna/tuna_api/internal/dto"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/repo"
)

type IAccountSvc interface {
	LoginUser(u dto.User)
}

type accountSvc struct {
	r repo.IAccountRepo
}

func NewAccountSvc(repo repo.IAccountRepo) accountSvc {
	return accountSvc{
		repo,
	}
}

func (as *accountSvc) LoginUser(u dto.User) {
	as.r.InsertUser()
}

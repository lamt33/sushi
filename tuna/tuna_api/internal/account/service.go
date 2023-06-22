package account

import (
	"github.com/lamt3/sushi/tuna/common/cache"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/dto"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/repo"
	uuid "github.com/satori/go.uuid"
)

type IAccountSvc interface {
	LoginUser(u dto.User)
	GetUser(userID uuid.UUID) (dto.User, error)
}

type accountSvc struct {
	r     repo.IAccountRepo
	cache cache.ICache
}

func NewAccountSvc(repo repo.IAccountRepo, cache cache.ICache) accountSvc {
	return accountSvc{
		repo,
		cache,
	}
}

func (as *accountSvc) LoginUser(u dto.User) {
	as.r.InsertUser()
	as.cache.Set(u.UserID.String(), u, 0)
}

func (as *accountSvc) GetUser(userID uuid.UUID) (dto.User, error) {
	v := dto.User{}
	err := as.cache.Get(userID.String(), &v)
	return v, err
}

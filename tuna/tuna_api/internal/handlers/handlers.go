package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lamt3/sushi/tuna/common/web"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/account"
	"github.com/lamt3/sushi/tuna/tuna_api/internal/dto"
)

type AccountHandler struct {
	acctSvc account.IAccountSvc
}

func NewAccountHandler(acctSvc account.IAccountSvc) *AccountHandler {
	return &AccountHandler{
		acctSvc: acctSvc,
	}
}

func (ah *AccountHandler) LoginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body := &dto.User{}
	err := web.ReadBody(r.Body, body)
	if err != nil {
		web.GenerateError(w, "invalid crmUser in request", 400)
		return
	}

	ah.acctSvc.LoginUser(*body)
	web.GenerateSuccess(w, "ok")
}

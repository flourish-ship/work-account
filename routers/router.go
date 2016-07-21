package routers

import (
	"github.com/flourish-ship/work-account/idao"
	"github.com/kataras/iris"
)

// AccountManager ...
type AccountManager struct {
	API *iris.Framework
	DAO idao.IDAO
}

// NewAccountManager ...
func NewAccountManager(daoImpl idao.IDAO) *AccountManager {
	return &AccountManager{
		API: iris.New(),
		DAO: daoImpl,
	}
}

func (am *AccountManager) initialize() {
	api := am.API
	api.Party("/v1/am")
	{
		api.Post("/login", am.Login)
	}
}

// Server ...
func (am *AccountManager) Server() {
	am.API.Listen(":3030")
}

package idao

import "github.com/flourish-ship/work-account/models"

// IDAO ...
type IDAO interface {
	IAccount
}

// IAccount ...
type IAccount interface {
	Login(username, password string) *models.Resp
}

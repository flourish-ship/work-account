package idao

import "github.com/flourish-ship/work-account/response"

// IDAO ...
type IDAO interface {
	IAccount
}

// IAccount ...
type IAccount interface {
	SignIn(username, password string) *response.Resp
}

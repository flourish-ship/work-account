package daomongo

import "github.com/flourish-ship/work-account/response"

// Login ...
func (dao *DAOMongo) Login(username, password string) *response.Resp {
	return &response.Resp{
		Code:    0,
		Message: "login succusess!",
	}
}

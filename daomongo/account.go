package daomongo

import "github.com/flourish-ship/work-account/response"

// SignIn ...
func (dao *DAOMongo) SignIn(username, password string) *response.Resp {
	return &response.Resp{
		Code:    0,
		Message: "Login succusess!",
	}
}

package daomongo

import "github.com/flourish-ship/work-account/models"

// Login ...
func (dao *DAOMongo) Login(username, password string) *models.Resp {
	return &models.Resp{
		Code:    0,
		Message: "login succusess!",
	}
}

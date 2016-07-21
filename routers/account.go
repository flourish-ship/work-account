package routers

import "github.com/kataras/iris"

// Login ...
func (am *AccountManager) Login(c *iris.Context) {
	var param struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}
	err := c.ReadForm(&param)
	if err != nil {

	}
	c.JSON(iris.StatusOK, am.DAO.Login(param.Username, param.Password))
}

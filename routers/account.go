package routers

import (
	"github.com/flourish-ship/work-account/idao"
	"github.com/flourish-ship/work-account/response"
	"github.com/kataras/iris"
)

// AccountRouter ...
type AccountRouter struct {
	R string
	//Redis  *redis.Database
	//API    *iris.Framework
	dao idao.IDAO
}

// Registe ...
func (ar *AccountRouter) Registe(am *AccountManager) {
	ar.dao = am.DAO
	api := am.API
	prefix := api.Party(ar.R)
	{
		prefix.Post("/signin", ar.SignIn)
	}
}

// SignIn ...
func (ar *AccountRouter) SignIn(c *iris.Context) {
	var param struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}
	err := c.ReadForm(&param)
	if err != nil {
		c.JSON(iris.StatusOK, response.RequestParamError.ErrReap())
		return
	}
	c.JSON(iris.StatusOK, ar.dao.SignIn(param.Username, param.Password))
}

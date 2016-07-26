package routers

import (
	"github.com/flourish-ship/work-account/idao"
	"github.com/flourish-ship/work-account/models"
	"github.com/flourish-ship/work-account/response"
	"github.com/kataras/iris"
)

// AccountRouter ...
type AccountRouter struct {
	R string
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
	param := models.SignInParam{}
	err := c.ReadForm(&param)
	if err != nil {
		c.JSON(iris.StatusOK, response.RequestParamError.ErrReap())
		return
	}

	result := ar.dao.SignIn(param.Username, param.Password)
	if result.Status == idao.NotFound {
		c.JSON(iris.StatusOK, &response.Resp{
			Code:    response.NotFound,
			Message: "Can't find this user",
		})
	} else if result.Status == idao.ValidationError {
		c.JSON(iris.StatusOK, &response.Resp{
			Code:    response.ValidationError,
			Message: "The account with this password was not found",
		})
	}
	if result.Status != idao.Succuess {
		return
	}
}

func (ar *AccountRouter) generateSessionAndSaveToken(c *iris.Context) {

}

package routers

import (
	"github.com/flourish-ship/work-account/auth/token"
	"github.com/flourish-ship/work-account/db"
	"github.com/flourish-ship/work-account/models"
	"github.com/flourish-ship/work-account/response"
	"github.com/kataras/iris"
)

// AccountRouter ...
type AccountRouter struct {
	R string
	//API    *iris.Framework
	dao *db.DAOMongo
}

// Registe ...
func (ar *AccountRouter) Registe(am *AccountManager) {
	ar.dao = am.DAO
	api := am.API
	prefix := api.Party(ar.R)
	{
		prefix.Post("/signin", ar.SignIn)
		prefix.Post("/signup", ar.SignUp)
		security := prefix.Party("/security", TokenAuthMiddleware)
		{
			security.Get("/logout", ar.SignOut)
		}
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
	// TODO data validate
	result := ar.dao.SignIn(param)
	if result.Status == db.NotFound {
		c.JSON(iris.StatusOK, &response.Resp{
			Code:    response.NotFound,
			Message: "Can't find this user",
		})
	} else if result.Status == db.ValidationError {
		c.JSON(iris.StatusOK, &response.Resp{
			Code:    response.ValidationError,
			Message: "The account with this password was not found",
		})
	}
	if result.Status != db.Succuess {
		return
	}
	user := result.Data.(models.User)
	tokenKey := token.GenerateAndSaveToken(c, user.Id.Hex())
	if tokenKey == "" {
		c.JSON(iris.StatusOK, response.TokenError.ErrReap())
		return
	}
	c.JSON(iris.StatusOK, &response.Resp{
		Code:    response.Succuess,
		Message: "Login succuess",
		Data:    tokenKey,
	})

}

// SignOut ...
func (ar *AccountRouter) SignOut(c *iris.Context) {
	token := c.Get("token").(token.Token)
	c.Session().Delete(token.Key)
	c.Session().Delete(token.UserID)
	c.JSON(iris.StatusOK, &response.Resp{
		Code:    response.Succuess,
		Message: "Logout succuess",
	})
}

// SignUp ...
func (ar *AccountRouter) SignUp(c *iris.Context) {
	user := models.User{}
	c.ReadJSON(&user)
	// TODO data validate
	result := ar.dao.SignUp(user)

}

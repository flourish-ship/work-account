package routers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/flourish-ship/work-account/auth/token"
	"github.com/flourish-ship/work-account/db"
	"github.com/flourish-ship/work-account/models"
	"github.com/flourish-ship/work-account/response"
	"github.com/kataras/iris"
	uuid "github.com/nu7hatch/gouuid"
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
	tokenKey, err := generateAndSaveToken(c, user.Id.Hex())
	if err != nil {
		c.JSON(iris.StatusOK, response.TokenError.ErrReap())
		return
	}
	c.JSON(iris.StatusOK, &response.Resp{
		Code:    response.Succuess,
		Message: "Login succuess",
		Data:    tokenKey,
	})

}

func generateAndSaveToken(c *iris.Context, userID string) (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	//uuid as token's key
	now := time.Now()
	tokenKey := strings.Replace(u.String(), "-", "", -1)
	token := token.Token{
		Key:         tokenKey,
		UserID:      userID,
		LastLoginAt: now.Unix(),
		ExpireAt:    now.Add(TOKENEXPIRE).Unix(),
	}
	tokenInfo, err := json.Marshal(&token)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Println("tokenKey", tokenKey)
	fmt.Println("tokenInfo", string(tokenInfo))
	fmt.Println("userID", userID)
	c.Session().Set(tokenKey, tokenInfo)
	c.Session().Set(userID, tokenKey)
	return tokenKey, nil
}

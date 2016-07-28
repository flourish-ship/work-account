package token

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/flourish-ship/work-account/response"
	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
)

const (
	// TOKENEXPIRE ...
	TOKENEXPIRE = time.Hour * 24 * 7
)

// Token ...
type Token struct {
	Key      string
	UserID   string
	ExpireAt int64
}

func GenerateAndSaveToken(c *iris.Context, userID string) string {
	u := uuid.NewV4()
	//uuid as token's key
	tokenKey := strings.Replace(u.String(), "-", "", -1)
	token := Token{
		Key:      tokenKey,
		UserID:   userID,
		ExpireAt: time.Now().Add(TOKENEXPIRE).Unix(),
	}
	tokenInfo, _ := json.Marshal(&token)
	//remove old token if exist
	oldTokenKey := c.Session().GetString(userID)
	if oldTokenKey != "" {
		c.Session().Delete(oldTokenKey)
	}
	c.Session().Set(tokenKey, tokenInfo)
	c.Session().Set(userID, tokenKey)
	//c.Session().VisitAll(func(k string, v interface{}) {
	//fmt.Println("k:", k)
	//fmt.Println("v:", v)
	//})
	return tokenKey
}

func TokenAuth(c *iris.Context, tokenKey string) *response.Resp {
	tokenInfo := c.Session().Get(tokenKey)
	if tokenInfo == nil {
		return response.NotFoundToken.ErrReap()
	}
	token := Token{}
	if json.Unmarshal(tokenInfo.([]byte), &token) != nil {
		//remove this token
		c.Session().Delete(tokenKey)
		return response.TokenAuthError.ErrReap()
	}
	if time.Now().Unix() > token.ExpireAt {
		return response.TokenExpire.ErrReap()
	}
	if c.Session().GetString(token.UserID) != tokenKey {
		return response.TokenInvalid.ErrReap()
	}
	//save in context
	c.Set("token", token)
	return nil
}

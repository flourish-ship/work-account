package routers

import (
	"strings"

	"github.com/flourish-ship/work-account/auth/token"
	"github.com/flourish-ship/work-account/response"
	"github.com/kataras/iris"
)

const (
	TOKENNAME = "WORK-TOKEN"
)

func TokenAuthMiddleware(c *iris.Context) {
	key := strings.TrimSpace(c.RequestHeader(TOKENNAME))
	if key == "" {
		c.JSON(iris.StatusOK, response.NoToken.ErrReap())
		return
	}
	//token auth
	resp := token.TokenAuth(c, key)
	if resp != nil {
		c.JSON(iris.StatusOK, resp)
		return
	}
	c.Next()
}

package daomongo

import (
	"github.com/flourish-ship/work-account/models"
	"github.com/flourish-ship/work-account/response"
	"gopkg.in/mgo.v2/bson"
)

// SignIn ...
func (dao *DAOMongo) SignIn(username, password string) *response.Resp {
	user := models.User{}
	resp := &response.Resp{}
	dao.db.C("users").Find(bson.M{"username": username}).One(&user)
	if user.Id == "" {
		resp.Code = int(response.NotFoundError)
		resp.Message = "Can't find this user!"
		return resp
	}
	if bcrypt.CompareHashAndPassword(password, []byte(pass)) != nil {
		return &models.ResponseExt{Code: "01", Message: "密码错误,登陆失败!"}
	}
}

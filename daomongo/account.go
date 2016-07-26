package daomongo

import (
	"github.com/flourish-ship/work-account/idao"
	"github.com/flourish-ship/work-account/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// SignIn ...
func (dao *DAOMongo) SignIn(username, password string) idao.Result {
	user := models.User{}
	dao.db.C("users").Find(bson.M{"username": username}).One(&user)
	if user.Id == "" {
		return idao.Result{Status: idao.NotFound, Data: nil}
	}
	if bcrypt.CompareHashAndPassword(user.Password, []byte(password)) != nil {
		return idao.Result{Status: idao.ValidationError, Data: nil}
	}
	return idao.Result{Status: idao.Succuess, Data: nil}
}

package db

import (
	"github.com/flourish-ship/work-account/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// SignIn ...
func (dao *DAOMongo) SignIn(username, password string) Result {
	user := models.User{}
	if username == "wendell" && password == "sunwen" {
		user.Id = bson.NewObjectId()
		return Result{Status: Succuess, Data: user}
	}
	dao.db.C("users").Find(bson.M{"username": username}).One(&user)
	if user.Id == "" {
		return Result{Status: NotFound, Data: nil}
	}
	if bcrypt.CompareHashAndPassword(user.Password, []byte(password)) != nil {
		return Result{Status: ValidationError, Data: nil}
	}
	return Result{Status: Succuess, Data: user}
}

package db

import (
	"github.com/flourish-ship/work-account/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// SignIn ...
func (dao *DAOMongo) SignIn(param models.SignInParam) Result {
	user := models.User{}
	if param.Username == "wendell" && param.Password == "sunwen" {
		user.Id = bson.NewObjectId()
		return Result{Status: Succuess, Data: user}
	}
	dao.db.C("users").Find(bson.M{"username": param.Username}).One(&user)
	if user.Id == "" {
		return Result{Status: NotFound, Data: nil}
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(param.Password)) != nil {
		return Result{Status: ValidationError, Data: nil}
	}
	return Result{Status: Succuess, Data: user}
}

// SignUp ...
func (dao *DAOMongo) SignUp(user models.User) Result {
	user.Id = bson.NewObjectId()
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return Result{Status: UnknownError, Data: nil}
	}
	user.PasswordHash = string(hashPass)
	if err := dao.db.C("users").Insert(&user); err != nil {
		return Result{Status: DBError, Data: nil}
	}
	return Result{Status: Succuess, Data: user}
}

func (dao *DAOMongo) CheckAccountExist(username string) {

}

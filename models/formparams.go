package models

type SignInParam struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type SignUpParam struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

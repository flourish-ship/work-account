package idao

// IDAO ...
type IDAO interface {
	IAccount
}

// IAccount ...
type IAccount interface {
	SignIn(username, password string) Result
}

type Result struct {
	Status
	Data interface{}
}

type Status int

const (
	Succuess Status = iota
	NotFound
	ValidationError
)

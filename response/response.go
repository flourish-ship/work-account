package response

type resultCode int

// Resp ...
type Resp struct {
	Code    resultCode  `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	Succuess resultCode = iota
	//RequestParamError ...
	NotFound
	ValidationError
	RequestParamError
	TokenError
	NoToken
	NotFoundToken
	TokenExpire
	TokenInvalid
	TokenAuthError
)

func (rc resultCode) ErrReap() *Resp {
	r := &Resp{Code: rc}
	switch rc {
	case RequestParamError:
		r.Message = "Request param error,please confirm and re submit"
	case TokenError:
		r.Message = "Error occurred when get/set token"
	case NoToken:
		r.Message = "Token not found from Header,not allowed access"
	case NotFoundToken:
		r.Message = "Token not found from Session,not allowed access"
	case TokenAuthError:
		r.Message = "Token auth failed,please re login"
	case TokenExpire:
		r.Message = "Token already expired,please re login"
	case TokenInvalid:
		r.Message = "Token already invalid,please re login"
	default:
		r.Message = "Unknown error,sorry"
	}
	return r
}

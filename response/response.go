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
	RequestParamError
	NotFound
	ValidationError
	TokenError
)

func (rc resultCode) ErrReap() *Resp {
	r := &Resp{Code: rc}
	switch rc {
	case RequestParamError:
		r.Message = "Request param error,please confirm and re submit!"
	case TokenError:
		r.Message = "Error occurred when get/set token!"
	default:
		r.Message = "Unknown error,sorry!"
	}
	return r
}

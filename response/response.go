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
)

func (rc resultCode) ErrReap() *Resp {
	r := &Resp{Code: rc}
	switch rc {
	case RequestParamError:
		r.Message = "Request param error,please confirm and re submit!"
	default:
		r.Message = "Unknown error,sorry!"
	}
	return r
}

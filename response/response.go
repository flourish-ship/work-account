package response

type errorType int

// Resp ...
type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	_ errorType = iota
	//RequestParamError ...
	RequestParamError
	NotFoundError
)

func (ey errorType) ErrReap() *Resp {
	r := &Resp{Code: int(ey)}
	switch ey {
	case RequestParamError:
		r.Message = "Request param error,please confirm and re submit!"
	default:
		r.Message = "Unknown error,sorry!"
	}
	return r
}

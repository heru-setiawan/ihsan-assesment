package payloads

import "assesment/pkg/exceptions"

type ResponseDefault struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func (res *ResponseDefault) ParseFromException(err error) (code int) {
	switch e := err.(type) {
	case exceptions.Validation:
		code = e.Code
		res.Message = e.Message
	case exceptions.Database:
		code = e.Code
		res.Message = e.Message
	default:
		code = 500
		res.Message = "internal server error"
	}

	return
}

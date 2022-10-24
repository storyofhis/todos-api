package common

import (
	"bytes"
	"encoding/json"
)

type Response struct {
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data any) Response {
	res := Response{
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err error) Response {
	res := Response{
		Message: message,
		Errors:  err.Error(),
		Data:    nil,
	}
	return res
}

func MapToStruct(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}

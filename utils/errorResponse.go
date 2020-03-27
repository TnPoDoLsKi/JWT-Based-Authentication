package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonError struct {
	Status int    `json:"status,omitempty"`
	Code   int    `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func (je JsonError) Error() []byte {
	out, _ := json.Marshal(je)
	return out
}

func ErrorResponse(w http.ResponseWriter, status int, code int, detail error) {
	js := JsonError{
		Status: status,
		Code:   code,
		Detail: detail.Error(),
	}

	if status >= 500 {
		log.Print(detail)
	}
	w.WriteHeader(status)
	w.Write(js.Error())
}

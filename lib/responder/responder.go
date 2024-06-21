package responder

import (
	"net/http"
)

type Responder interface {
	OutputJson(writer http.ResponseWriter, jsonData []byte)
	OutputMessage(writer http.ResponseWriter, status int, messages ...string)
}

type Respond struct {
}

func New() Responder {
	return &Respond{}
}

func (r *Respond) OutputJson(writer http.ResponseWriter, jsonData []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonData)
}

func (r *Respond) OutputMessage(writer http.ResponseWriter, status int, messages ...string) {
	writer.WriteHeader(status)
	for _, message := range messages {
		writer.Write([]byte(message))
	}
}

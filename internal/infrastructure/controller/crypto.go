package controller

import (
	"cryptoStringRedisCache/internal/infrastructure/dto"
	"cryptoStringRedisCache/internal/infrastructure/service"
	"cryptoStringRedisCache/lib/responder"
	"encoding/json"
	"log"
	"net/http"
)

type Crypter interface {
	Crypto(writer http.ResponseWriter, request *http.Request)
}

type Crypto struct {
	responder responder.Responder
	service   service.Crypter
}

func NewCrypto(responder responder.Responder, service service.Crypter) Crypter {
	return &Crypto{responder: responder, service: service}
}

func (c *Crypto) Crypto(writer http.ResponseWriter, request *http.Request) {
	var requestDTO dto.Request
	if err := json.NewDecoder(request.Body).Decode(&requestDTO); err != nil {
		c.responder.OutputMessage(writer, http.StatusBadRequest, err.Error())
		return
	}

	md5, err := c.service.MD5(requestDTO.Query)
	if err != nil {
		c.responder.OutputMessage(writer, http.StatusInternalServerError, err.Error())
		return
	}
	sha256, err := c.service.SHA256(requestDTO.Query)
	if err != nil {
		c.responder.OutputMessage(writer, http.StatusInternalServerError, err.Error())
		return
	}
	result := dto.Response{MD5: md5, SHA256: sha256}
	resultJson, err := json.Marshal(&result)
	if err != nil {
		log.Printf("err response to json: %s\n", err)
		c.responder.OutputMessage(writer, http.StatusBadRequest)
		return
	}
	c.responder.OutputJson(writer, resultJson)
}

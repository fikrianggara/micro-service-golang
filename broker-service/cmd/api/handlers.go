package main

import (
	"encoding/json"
	"net/http"
)

//json:error, json:message, json:data omitempty
//merupakan 'literal tag' yang berfungsi sebagai atribut dari semua field
//sesuai field tempat deklarasinya, di web bisa berguna untuk json, xml, protobuf, dst
type jsonResponse struct {
	Error   bool   `json:error`
	Message string `json:message`
	Data    any    `json:data, omitempty`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "hi from broker service",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}

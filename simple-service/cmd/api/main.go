package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:error`
	Message string `json:message`
	Data    any    `json:data, omitempty`
}

func main() {
	AppPort := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "hello from simple service")
	})
	http.HandleFunc("/endpoint2", func(w http.ResponseWriter, _ *http.Request) {
		payload := jsonResponse{
			Error:   false,
			Message: "hi from simple service",
		}

		out, _ := json.MarshalIndent(payload, "", "\t")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(out)
	})
	log.Fatal(http.ListenAndServe(":"+AppPort, nil))
}

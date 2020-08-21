package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func _RespondWithERROR(w http.ResponseWriter, code int, msg string) {
	_RespondWithJSON(w, code, msg)
}

func _RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

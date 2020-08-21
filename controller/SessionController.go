package controller

import (
	"encoding/json"
	"net/http"

	"github.com/openpracticetool/maestro/model"
)

//SaveSession save session in database
func SaveSession(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var session model.SessionModel

	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		_RespondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {}

func UpdateSesion(w http.ResponseWriter, r *http.Request) {}

func FindSessionByID(w http.ResponseWriter, r *http.Request) {}

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/openpracticetool/maestro/model"
	"github.com/openpracticetool/maestro/repository"
)

var sr repository.SessionRepository

//SaveSession save session in database
func SaveSession(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var session model.Session

	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		_RespondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	sr.SaveSession(session)

	//Return the status
	_RespondWithJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

//DeleteSession delete session of database
func DeleteSession(w http.ResponseWriter, r *http.Request) {

}

//UpdateSesion update session in database
func UpdateSesion(w http.ResponseWriter, r *http.Request) {

}

//FindSessionByID find session by id
func FindSessionByID(w http.ResponseWriter, r *http.Request) {

}

//FindSessionByUserCreation find session by user creation
func FindSessionByUserCreation(w http.ResponseWriter, r *http.Request) {

}

//FindSessionByWorkspaceID find session by workspace id
func FindSessionByWorkspaceID(w http.ResponseWriter, r *http.Request) {

}

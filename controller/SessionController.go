package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"github.com/openpracticetool/maestro/model"
	"github.com/openpracticetool/maestro/repository"
	"github.com/openpracticetool/maestro/validators"
)

// SessionController ::: struct
type SessionController struct {
	db      *gorm.DB
	message []string
}

// NewSessionController ::: resturns a new SessionController
func NewSessionController(db *gorm.DB) *SessionController {
	return &SessionController{
		db: db,
	}
}

//SaveSession ::: save session in database
func (sc *SessionController) SaveSession(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var session model.Session

	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		// print the log in console
		log.Println(err)

		_RespondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	// initiliza the session repository
	sr := repository.NewSessionRespository(sc.db)

	if err := validators.ValidateStruct(session); err != nil {
		for i, err := range err.(validator.ValidationErrors) {

			switch {
			case err.Field() == "Description":
				sc.message[i] = "The Description should be contain min 10 charactres and max 255."
			case err.Field() == "Name":
				sc.message[i] = "The Name should be contain min 10 characteres and max 50."
			case err.Field() == "IDWorkspace":
				sc.message[i] = "The IDWorkspace can't be nullabel."
			}
		}

		// print the erro in console
		log.Println(sc.message)
		// Send the error to requester
		_RespondWithArrayERROR(w, http.StatusBadRequest, sc.message)
	}

	// set date
	session.CreatedAt = time.Now()

	// save the session in database
	session, err := sr.SaveSession(session)

	// check if the procedure occurs with errors
	if err != nil {
		// print the message in console
		log.Println(err)
		// send the error to requester
		_RespondWithERROR(w, http.StatusInternalServerError, "Error to save Session in database.")
	}
	// send the session created to resquester
	_RespondWithJSON(w, http.StatusOK, session)
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

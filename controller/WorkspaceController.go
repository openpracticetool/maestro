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

// WorkspaceController ::: struct
type WorkspaceController struct {
	db      *gorm.DB
	message []string
}

// NewWorkspaceController ::: returns a new WorkspaceController
func NewWorkspaceController(db *gorm.DB) *WorkspaceController {
	return &WorkspaceController{
		db: db,
	}
}

//SaveWorkspace save workspace in database
func (wc *WorkspaceController) SaveWorkspace(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var workspace model.Workspace

	if err := json.NewDecoder(r.Body).Decode(&workspace); err != nil {
		// print the log in console
		log.Println(err)

		_RespondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
	// initiliza the workspace repository
	wr := repository.NewWorkspaceRepository(wc.db)

	if err := validators.ValidateStruct(workspace); err != nil {
		for i, err := range err.(validator.ValidationErrors) {

			switch {
			case err.Field() == "Description":
				wc.message[i] = "The Description should be contain min 10 charactres and max 255."
			case err.Field() == "Name":
				wc.message[i] = "The Name should be contain min 10 characteres and max 50."
			case err.Field() == "IDWorkspace":
				wc.message[i] = "The IDWorkspace can't be nullabel."
			}
		}

		// print the erro in console
		log.Println(wc.message)
		// Send the error to requester
		_RespondWithArrayERROR(w, http.StatusBadRequest, wc.message)
	}

	// set date
	workspace.CreatedAt = time.Now()

	// save the workspace in database
	workspace, err := wr.SaveWorkspace(workspace)

	// check if the procedure occurs with errors
	if err != nil {
		// print the message in console
		log.Println(err)
		// send the error to requester
		_RespondWithERROR(w, http.StatusInternalServerError, "Error to save Workspace in database.")
	}
	// send the workspace created to resquester
	_RespondWithJSON(w, http.StatusOK, workspace)
}

//DeleteWorkspace delete workspace of database
func (wc *WorkspaceController) DeleteWorkspace(w http.ResponseWriter, r *http.Request) {}

//UpdateWorkspace update workspace in database
func (wc *WorkspaceController) UpdateWorkspace(w http.ResponseWriter, r *http.Request) {}

//FindWorkspaceByID find workspace by id
func (wc *WorkspaceController) FindWorkspaceByID(w http.ResponseWriter, r *http.Request) {}

//FindWorkspaceByUserCreation find workspace by user creation
func FindWorkspaceByUserCreation(w http.ResponseWriter, r *http.Request) {}

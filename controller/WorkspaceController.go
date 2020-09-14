package controller

import (
	"encoding/json"
	"net/http"

	"github.com/openpracticetool/maestro/model"
)

//SaveWorkspace save workspace in database
func SaveWorkspace(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var workspace model.Workspace

	if err := json.NewDecoder(r.Body).Decode(&workspace); err != nil {
		_RespondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
}

//DeleteWorkspace delete workspace of database
func DeleteWorkspace(w http.ResponseWriter, r *http.Request) {}

//UpdateWorkspace update workspace in database
func UpdateWorkspace(w http.ResponseWriter, r *http.Request) {}

//FindWorkspaceByID find workspace by id
func FindWorkspaceByID(w http.ResponseWriter, r *http.Request) {}

//FindWorkspaceByUserCreation find workspace by user creation
func FindWorkspaceByUserCreation(w http.ResponseWriter, r *http.Request) {}

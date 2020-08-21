package controller

import (
	"encoding/json"
	"net/http"

	"github.com/openpracticetool/maestro/model"
)

//WorkspaceModel save workspace in database
func SaveWorkspace(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var workspace model.WorkspaceModel

	if err := json.NewDecoder(r.Body).Decode(&workspace); err != nil {
		_RespondWithERROR(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
}

func DeleteWorkspace(w http.ResponseWriter, r *http.Request) {}

func UpdateWorkspace(w http.ResponseWriter, r *http.Request) {}

func FindWorkspaceByID(w http.ResponseWriter, r *http.Request) {}

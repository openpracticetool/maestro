package test

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/openpracticetool/maestro/model"
	"github.com/openpracticetool/maestro/repository"
)

var database = repository.Database{}
var workspace = model.Workspace{}

var db *gorm.DB

func init() {
	database.Server = "postgresql://root@0.0.0.0:26257/opt?sslmode=disable"
	database.LogMode = false

	db = database.Connect()
}

func TestDBConnection(t *testing.T) {
	if db.Error != nil {
		t.Error("Error to connect in database")
	}
}

func TestPingDatabse(t *testing.T) {
	if err := db.DB().Ping(); err != nil {
		t.Error("Erro to ping database")
	}
}

func TestSaveWorkspace(t *testing.T) {
	// defines the workspace to save
	var model = model.Workspace{
		IDFacilitator: 123456789,
		Description:   "Este é um exmplo de teste para criação de workspace",
		Name:          "Pratica com Itau",
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
		UpdatedBY:     "lhsribas",
		CreatedBy:     "lhsribas",
	}

	// Initilize a repository with a connection of database
	wr := repository.NewWorkspaceRepository(db)
	_, err := wr.SaveWorkspace(model)

	if err != nil && workspace.ID > 0 {
		t.Fail()
	}
}

func TestUpdateWorkspace(t *testing.T) {
	var ID int
	// Initilize a repository with a connection of database
	wr := repository.NewWorkspaceRepository(db)

	// Find all workspaces in database
	workspaces, _ := wr.FindWorkspaceByCreatedBy("lhsribas")

	// Select the first workspace and get your ID
	for _, v := range workspaces {
		ID = v.ID
		break
	}
	// Find a workspace by a specific ID
	workspace, err := wr.FindWorkspaceByID(ID)

	workspace.Description = "Novo tema para discutir sobre a alteração das informações"
	workspace.Name = "Test Update"
	workspace.UpdatedAt = time.Now()
	workspace.UpdatedBY = "lhsribas2"

	response, err := wr.UpdateWorkspace(workspace)

	// Verify if the test are ok
	if err != nil || workspace.ID != ID {
		t.Fail()
	}

	// compare the fields with information returned of database
	if response.Name != workspace.Name || response.Description != workspace.Description || response.UpdatedBY != workspace.UpdatedBY {
		t.Fail()
	}
}

func TestFindWorkspaceByCreatedBy(t *testing.T) {
	// Initilize a repository with a connection of database
	wr := repository.NewWorkspaceRepository(db)

	workspaces, err := wr.FindWorkspaceByCreatedBy("lhsribas")

	if err != nil || len(workspaces) == 0 {
		t.Fail()
	}
}

func TestFindWorkspaceByID(t *testing.T) {
	var ID int
	// Initilize a repository with a connection of database
	wr := repository.NewWorkspaceRepository(db)

	// Find all workspaces in database
	workspaces, _ := wr.FindWorkspaceByCreatedBy("lhsribas")

	// Select the first workspace and get your ID
	for _, v := range workspaces {
		ID = v.ID
		break
	}
	// Find a workspace by a specific ID
	workspace, err := wr.FindWorkspaceByID(ID)

	// Verify if the test are ok
	if err != nil || workspace.ID != ID {
		t.Fail()
	}
}

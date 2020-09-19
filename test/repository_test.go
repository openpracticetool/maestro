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
var session = model.Session{}

var db *gorm.DB

func init() {
	database.Server = "postgresql://root@0.0.0.0:26257/opt?sslmode=disable"
	database.LogMode = false

	db = database.Connect()
}

// TestDBConnection ::: check if the connection return some error
func TestDBConnection(t *testing.T) {
	if db.Error != nil {
		t.Error("Error to connect in database")
	}
}

//TestPingDatabse ::: execute a ping in database to test the connection
func TestPingDatabse(t *testing.T) {
	if err := db.DB().Ping(); err != nil {
		t.Error("Erro to ping database")
	}
}

// TestSaveWorkspace save a workspace in database
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

	// Save the workspace in database
	workspace, err := wr.SaveWorkspace(model)

	// check if the procedure return some error
	if err != nil && workspace.ID > 0 {
		t.Fail()
	}
}

// TestUpdateWorkspace :::
func TestUpdateWorkspace(t *testing.T) {
	var ID = workspace.ID
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
	workspace, err := wr.FindWorkspaceByID(workspace.ID)

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

// TestFindWorkspaceByCreatedBy :::
func TestFindWorkspaceByCreatedBy(t *testing.T) {
	// Initilize a repository with a connection of database
	wr := repository.NewWorkspaceRepository(db)

	// find workspaces by user creation
	workspaces, err := wr.FindWorkspaceByCreatedBy("lhsribas")

	if err != nil || len(workspaces) == 0 {
		t.Fail()
	}
}

// TestFindWorkspaceByID :::
func TestFindWorkspaceByID(t *testing.T) {
	// Initilize a repository with a connection of database
	wr := repository.NewWorkspaceRepository(db)

	// Find all workspaces in database
	workspaces, _ := wr.FindWorkspaceByCreatedBy("lhsribas")

	// Select the first workspace and get your ID
	for _, v := range workspaces {
		// Find a workspace by a specific ID
		workspace, err := wr.FindWorkspaceByID(v.ID)

		// Verify if the test are ok
		if err != nil || workspace.ID != v.ID {
			t.Fail()
		}
		break
	}
}

// TestDeleteWorkspaceByID :::
func TestDeleteWorkspaceByID(t *testing.T) {
	// initialize a repository with database connection
	wr := repository.NewWorkspaceRepository(db)

	// find the workspace like user creation
	workspaces, err := wr.FindWorkspaceLikeCreatedBy("lhsribas")

	// check if the operation return some error
	if err != nil {
		t.Fail()
	}

	// iterate the results and delete of database
	for _, v := range workspaces {
		// delete all workspaces created
		if err := wr.DeleteWorkspace(v.ID); err != nil {
			t.Fail()
		}
	}
}

// TestSaveSession :::
func TestSaveSession(t *testing.T) {
	// initialize a repository with database connection
	sr := repository.NewSessionRespository(db)

	var model = model.Session{
		IDWorkspace: 123456789,
		Description: "Este é um exmplo de teste para criação de uma seção",
		Name:        "Lean Coffee table",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedBY:   "lhsribas",
		CreatedBy:   "lhsribas",
	}

	// save the session in database
	_, err := sr.SaveSession(model)

	// check if returns some error and fail the test
	if err != nil && workspace.ID > 0 {
		t.Fail()
	}
}

// TestUpdateSession :::
func TestUpdateSession(t *testing.T) {
	// initialize a repository with database connection
	sr := repository.NewSessionRespository(db)

	// find sessions by id of workspace
	sessions, err := sr.FindSessionByWorkspaceID(123456789)

	// check if return some error
	if err != nil {
		t.Fail()
	}

	// iterate the sessions and select the first
	for _, v := range sessions {
		session = v
		break
	}

	session.Description = "Novo tema para discutir sobre a alteração das informações"
	session.Name = "Test Update"
	session.UpdatedAt = time.Now()
	session.UpdatedBY = "lhsribas2"

	// update the session
	response, err := sr.UpdateSession(session)

	// compare the fields with information returned of database
	if err != nil ||
		session.ID != response.ID ||
		response.Name != session.Name ||
		response.Description != session.Description ||
		response.UpdatedBY != session.UpdatedBY {

		t.Fail()
	}
}

// TestFindSessionByID :::
func TestFindSessionByID(t *testing.T) {
	// initialize a repository with database connection
	sr := repository.NewSessionRespository(db)

	// find session by id of workspace
	sessions, err := sr.FindSessionByWorkspaceID(123456789)

	// check if the procedure return some error
	if err != nil {
		t.Fail()
	}

	// iterate the sessions and select the first
	for _, v := range sessions {
		// find session by id
		session, err := sr.FindSessionByID(v.ID)

		// check if the procedure return some error
		if err != nil || session.ID <= 0 {
			t.Fail()
		}

		break
	}
}

// TestFindSessionByWorkspaceID :::
func TestFindSessionByWorkspaceID(t *testing.T) {
	// initialize a repository with database connection
	sr := repository.NewSessionRespository(db)

	// find the sessions by id of workspace
	sessions, err := sr.FindSessionByWorkspaceID(123456789)

	if err != nil || len(sessions) <= 0 {
		t.Fail()
	}
}

// TestDeleteSessionID :::
func TestDeleteSessionID(t *testing.T) {
	// initialize a repository with database connection
	sr := repository.NewSessionRespository(db)

	// find the sessions by id of workspace
	sessions, err := sr.FindSessionByWorkspaceID(123456789)

	// check if the procedure return some error
	if err != nil {
		t.Fail()
	}

	// iterate the session
	for _, v := range sessions {
		// delete all sessions created
		if err := sr.DeleteSession(v.ID); err != nil {
			t.Fail()
		}
	}
}

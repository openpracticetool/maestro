package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/openpracticetool/maestro/converter"
	"github.com/openpracticetool/maestro/model"
)

// WorkspaceRepository struct data
type WorkspaceRepository struct {
	db *gorm.DB
}

// NewWorkspaceRepository returns a new WorkspaceRepository
func NewWorkspaceRepository(db *gorm.DB) *WorkspaceRepository {
	return &WorkspaceRepository{
		db: db,
	}
}

// SaveWorkspace save the Workspace data in database
func (wr *WorkspaceRepository) SaveWorkspace(workspace model.Workspace) (model.Workspace, error) {
	db := wr.db.Save(&workspace)

	if db.Error != nil {
		return model.Workspace{}, db.Error
	}

	err := converter.ConverterInterfaceTOStruct(db.Value, &workspace)
	return workspace, err
}

// UpdateWorkspace update the workspace data in database
func (wr *WorkspaceRepository) UpdateWorkspace(workspace model.Workspace) (model.Workspace, error) {
	// updates the conlumns od database
	db := wr.db.Model(&workspace).Where("id = ? ", workspace.ID).UpdateColumns(model.Workspace{Description: workspace.Description, Name: workspace.Name, UpdatedBY: workspace.UpdatedBY, UpdatedAt: time.Now()})

	// checks if database return errror
	if db.Error != nil {
		return model.Workspace{}, db.Error
	}

	// convert the interface response in struct
	err := converter.ConverterInterfaceTOStruct(db.Value, &workspace)

	return workspace, err
}

// DeleteWorkspace delete the workspace data in database
func (wr *WorkspaceRepository) DeleteWorkspace(ID int) error {
	var workspace = model.Workspace{}
	// delete the workspace in database
	db := wr.db.Delete(&workspace, ID)

	return db.Error
}

// FindWorkspaceByID find the workspace by ID in database
func (wr *WorkspaceRepository) FindWorkspaceByID(ID int) (model.Workspace, error) {
	var workspace = model.Workspace{}

	// convert the interface response in struct
	err := converter.ConverterInterfaceTOStruct(wr.db.First(&workspace, ID).Value, &workspace)

	return workspace, err
}

// FindWorkspaceByCreatedBy find the workspaces by user creation
func (wr *WorkspaceRepository) FindWorkspaceByCreatedBy(createdBy string) ([]model.Workspace, error) {
	// initializes an array
	var workspaces = []model.Workspace{}
	// consulting in the database with the createdby parameter
	db := wr.db.Where("created_by = ?", createdBy).Find(&workspaces)

	// check if the db return error
	if db.Error != nil {
		return workspaces, db.Error
	}

	// convert the interface response in struct
	err := converter.ConverterInterfaceTOStruct(db.Value, &workspaces)

	return workspaces, err
}

package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/openpracticetool/maestro/converter"
	"github.com/openpracticetool/maestro/model"
)

// SessionRepository ::: struct
type SessionRepository struct {
	db *gorm.DB
}

// NewSessionRespository ::: returns a new SessionRepository
func NewSessionRespository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

// SaveSession ::: save the Session data in database
func (sr *SessionRepository) SaveSession(session model.Session) (model.Session, error) {
	// save the object session in database
	db := sr.db.Save(&session)

	// check with returned some error of procedure
	if db.Error != nil {
		return model.Session{}, db.Error
	}

	// convert the value returned to db for a session object
	err := converter.ConverterInterfaceTOStruct(db.Value, &session)
	return session, err
}

// UpdateSession ::: update the Session data in database
func (sr *SessionRepository) UpdateSession(session model.Session) (model.Session, error) {
	// updates the columns
	db := sr.db.Model(&session).Where("id = ?", session.ID).UpdateColumns(model.Session{Name: session.Name, Description: session.Description, UpdatedBY: session.UpdatedBY, UpdatedAt: time.Now()})

	// check with returned some error of procedure
	if db.Error != nil {
		return model.Session{}, db.Error
	}

	// convert the value returned to db for a session object
	err := converter.ConverterInterfaceTOStruct(db.Value, &session)
	return session, err

}

// DeleteSession ::: delete the Session data in database
func (sr *SessionRepository) DeleteSession(ID int) error {
	var session = model.Session{}
	// delete a session of database
	db := sr.db.Delete(&session, ID)

	return db.Error
}

// FindSessionByID ::: find the Session by ID in database
func (sr *SessionRepository) FindSessionByID(ID int) (model.Session, error) {
	// create a variable session
	var session = model.Session{}

	// find the session by ID
	db := sr.db.First(&session, ID)

	// check with returned some error of procedure
	if db.Error != nil {
		return model.Session{}, db.Error
	}

	// convert the value returned to db for a session object
	err := converter.ConverterInterfaceTOStruct(db.Value, &session)
	return session, err
}

// FindSessionByWorkspaceID ::: Find the Session by ID of Workspace
func (sr *SessionRepository) FindSessionByWorkspaceID(workspaceID int) ([]model.Session, error) {
	// create an array of sessions
	var sessions = []model.Session{}

	// find a session by ID of an workspace
	db := sr.db.Where("id_workspace = ?", workspaceID).Find(&sessions)

	// check with returned some error of procedure
	if db.Error != nil {
		return []model.Session{}, db.Error
	}

	// convert the value returned to db for a session object
	err := converter.ConverterInterfaceTOStruct(db.Value, &sessions)
	return sessions, err
}

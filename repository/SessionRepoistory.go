package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/openpracticetool/maestro/model"
)

// SessionRepository struct
type SessionRepository struct {
	db *gorm.DB
}

// NewSessionRespository returns a new SessionRepository
func NewSessionRespository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

// SaveSession save the Session data in database
func (sr *SessionRepository) SaveSession(Session model.Session) {
	//GetConnection().Create(&Session)
}

// UpdateSession update the Session data in database
func (sr *SessionRepository) UpdateSession() {

}

// DeleteSession delete the Session data in database
func (sr *SessionRepository) DeleteSession() {

}

// FindSessionByID find the Session by ID in database
func (sr *SessionRepository) FindSessionByID() {

}

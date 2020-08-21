package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/openpracticetool/maestro/model"
)

var db *gorm.DB

//Database model to pass params
type Database struct {
	Server  string
	LogMode bool
}

func (dbConnctionParams *Database) Connect() {
	db, err := gorm.Open("postgres", dbConnctionParams.Server)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Enable LogMode to collect infos
	db.LogMode(dbConnctionParams.LogMode)

	//Generate tables
	generateDB()
}

//Return the connection with database
func (dbConnctionParams *Database) getConnection() *gorm.DB {
	return db
}

//Generates Tabales in Database
func generateDB() {
	// Automatically create the "tables" based on the Models
	// model.
	db.AutoMigrate(&model.SessionModel{})
	db.AutoMigrate(&model.WorkspaceModel{})
}

package repository

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/openpracticetool/maestro/model"
)

var db *gorm.DB

//Database model to pass params
type Database struct {
	Server  string
	LogMode bool
}

// Functions of type `txnFunc` are passed as arguments to our
// `runTransaction` wrapper that handles transaction retries for us
// (see implementation below).
type txnFunc func(*gorm.DB) error

// Connect this func connect with database
func (d *Database) Connect() *gorm.DB {

	db, err := gorm.Open("postgres", d.Server)

	if err != nil {
		panic(err.Error())
	}

	// Enable LogMode to collect infos
	db.LogMode(d.LogMode)

	// Generate tables
	generateDB(db)

	return db
}

// This function is used for testing the transaction retry loop.  It
// can be deleted from production code.
var forceRetryLoop txnFunc = func(db *gorm.DB) error {

	// The first statement in a transaction can be retried transparently
	// on the server, so we need to add a placeholder statement so that our
	// force_retry statement isn't the first one.
	if err := db.Exec("SELECT now()").Error; err != nil {
		return err
	}

	// Used to force a transaction retry.
	if err := db.Exec("SELECT crdb_internal.force_retry('1s'::INTERVAL)").Error; err != nil {
		return err
	}
	return nil
}

// Wrapper for a transaction.  This automatically re-calls `fn` with
// the open transaction as an argument as long as the database server
// asks for the transaction to be retried.
func runTransaction(db *gorm.DB, fn txnFunc) error {
	var maxRetries = 3
	for retries := 0; retries <= maxRetries; retries++ {
		if retries == maxRetries {
			return fmt.Errorf("hit max of %d retries, aborting", retries)
		}
		txn := db.Begin()
		if err := fn(txn); err != nil {
			// We need to cast GORM's db.Error to *pq.Error so we can
			// detect the Postgres transaction retry error code and
			// handle retries appropriately.
			pqErr := err.(*pq.Error)
			if pqErr.Code == "40001" {
				// Since this is a transaction retry error, we
				// ROLLBACK the transaction and sleep a little before
				// trying again.  Each time through the loop we sleep
				// for a little longer than the last time
				// (A.K.A. exponential backoff).
				txn.Rollback()
				var sleepMs = math.Pow(2, float64(retries)) * 100 * (rand.Float64() + 0.5)
				fmt.Printf("Hit 40001 transaction retry error, sleeping %f milliseconds\n", sleepMs)
				time.Sleep(time.Millisecond * time.Duration(sleepMs))
			} else {
				// If it's not a retry error, it's some other sort of
				// DB interaction error that needs to be handled by
				// the caller.
				return err
			}
		} else {
			// All went well, so we try to commit and break out of the
			// retry loop if possible.
			if err := txn.Commit().Error; err != nil {
				pqErr := err.(*pq.Error)
				if pqErr.Code == "40001" {
					// However, our attempt to COMMIT could also
					// result in a retry error, in which case we
					// continue back through the loop and try again.
					continue
				} else {
					// If it's not a retry error, it's some other sort
					// of DB interaction error that needs to be
					// handled by the caller.
					return err
				}
			}
			break
		}
	}
	return nil
}

// generates Tabales in Database
func generateDB(db *gorm.DB) {
	// Automatically create the "tables" based on the Models
	// model.
	db.AutoMigrate(&model.Session{})
	db.AutoMigrate(&model.Workspace{})
}

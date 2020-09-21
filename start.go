package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/openpracticetool/maestro/config"
	"github.com/openpracticetool/maestro/controller"
	"github.com/openpracticetool/maestro/repository"
)

var c = config.Config{}
var database = repository.Database{}

var db *gorm.DB

//Const
const (
	PORT = ":3000"
)

/********
 * Initilize the database
 ********/
func init() {
	c.Read()

	database.Server = c.Server
	database.LogMode = c.LogMode

	db = database.Connect()
}

/*******
 * This func up a server in a specific port
 *******/
func serverUP(router *mux.Router) {
	fmt.Printf("Server Running in port: %s", PORT)
	http.ListenAndServe(PORT, router)
}

/*******
 * This func create a router
 *******/
func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	addRouter(router)

	return router
}

/*******
 * This func add new routes to a router
 *******/
func addRouter(router *mux.Router) {

	sc := controller.NewSessionController(db)
	wc := controller.NewWorkspaceController(db)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	//Creates a subrouter
	subRouter := router.PathPrefix("/maestro/api").Subrouter()

	//Add subrouter workspace
	subRouter.HandleFunc("/v1/workspaces", wc.SaveWorkspace).Methods("POST")
	subRouter.HandleFunc("/v1/workspaces/createdby/{created_by}", wc.FindWorkspaceByCreatedBy).Methods("GET")
	subRouter.HandleFunc("/v1/workspaces/workspaceid/{id_workspace}", wc.FindWorkspaceByID).Methods("GET")

	//Add subrouter session
	subRouter.HandleFunc("/v1/sessions", sc.SaveSession).Methods("POST")
}

/*******
 * Running MicroService
 *******/
func main() {
	serverUP(newRouter())
}

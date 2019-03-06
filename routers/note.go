package routers

import (
	"github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//Notes Routes
func SetNoteRoutes(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()

	noteRouter.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	noteRouter.HandleFunc("/notes/{id}", controllers.UpdateNote).Methods("PUT")
	noteRouter.HandleFunc("/notes/{id}", controllers.GetNoteByID).Methods("GET")
	noteRouter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
	noteRouter.HandleFunc("notes/tasks/{id}", controllers.GetNotesByTask).Methods("GET")
	noteRouter.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRouter),
	))

	return router
}

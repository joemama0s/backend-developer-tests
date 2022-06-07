package router

import (
	"github.com/gorilla/mux"
	"github.com/joemama0s/backend-developer-tests/rest-service/pkg/models"
)

func NewPersonRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/people/{id}", models.FindPersonByID).Methods("GET")
	r.HandleFunc("/people", models.FindPeopleByName).Methods("GET").Queries("first_name", "{first_name}", "last_name", "{last_name}")
	r.HandleFunc("/people", models.FindPeopleByPhoneNumber).Methods("GET").Queries("phone_number", "{phone_number}")
	r.HandleFunc("/people", models.AllPeople).Methods("GET")

	return r
}

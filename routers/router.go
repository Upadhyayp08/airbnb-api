package router

import (
	"airbnb-api/handlers"

	"github.com/gorilla/mux"
)

// InitializeRouter sets up the router
func InitializeRouter() *mux.Router {
	r := mux.NewRouter()

	// Define the route for Airbnb data
	r.HandleFunc("/api/airbnb/{room_id}", handlers.GetAirbnbData).Methods("GET")

	return r
}

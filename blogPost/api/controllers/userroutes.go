package controllers

import (
	services "blogPost/api/services"
)

func (s *Server) initializeUserRoutes() {

	//Insert users details
	s.Router.HandleFunc("/users", services.FormatJSONResponse(s.CreateUser)).Methods("POST")
	//Get all the user details
	s.Router.HandleFunc("/users", services.FormatJSONResponse(s.GetUsers)).Methods("GET")
	//Get user details by user id
	s.Router.HandleFunc("/users/{id}", services.FormatJSONResponse(s.GetUser)).Methods("GET")
	//Update user details by user id by authentication
	s.Router.HandleFunc("/users/{id}", services.FormatJSONResponse(services.Authenticate(s.UpdateUser))).Methods("PUT")
	//Delete user by user id by authentication
	s.Router.HandleFunc("/users/{id}", services.Authenticate(s.DeleteUser)).Methods("DELETE")

}

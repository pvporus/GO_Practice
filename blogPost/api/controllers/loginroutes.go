package controllers

import (
	services "blogPost/api/services"
)

func (s *Server) initializeLoginRoutes() {

	s.Router.HandleFunc("/login", services.FormatJSONResponse(s.Login)).Methods("POST")

}

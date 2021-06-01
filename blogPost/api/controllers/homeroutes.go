package controllers

import (
	services "blogPost/api/services"
)

func (s *Server) initializeHomeRoutes() {
	//HOme route
	s.Router.HandleFunc("/", services.FormatJSONResponse(s.Home)).Methods("GET")

}

package controllers

import (
	services "blogPost/api/services"
)

func (s *Server) initializePostsRoutes() {

	//Insert posts details
	s.Router.HandleFunc("/posts", services.FormatJSONResponse(s.CreatePost)).Methods("POST")
	//Get all hte posts details
	s.Router.HandleFunc("/posts", services.FormatJSONResponse(s.GetPosts)).Methods("GET")
	//Get specific posts details by posts id
	s.Router.HandleFunc("/posts/{id}", services.FormatJSONResponse(s.GetPost)).Methods("GET")
	//Update post by authentication
	s.Router.HandleFunc("/posts/{id}", services.FormatJSONResponse(services.Authenticate(s.UpdatePost))).Methods("PUT")
	//Delete post by authentication
	s.Router.HandleFunc("/posts/{id}", services.Authenticate(s.DeletePost)).Methods("DELETE")

}

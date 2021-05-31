package controllers

import services "blogPost/api/services"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", services.FormatJSONResponse(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", services.FormatJSONResponse(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", services.FormatJSONResponse(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", services.FormatJSONResponse(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", services.FormatJSONResponse(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", services.FormatJSONResponse(services.Authenticate(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", services.Authenticate(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", services.FormatJSONResponse(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", services.FormatJSONResponse(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", services.FormatJSONResponse(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", services.FormatJSONResponse(services.Authenticate(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", services.Authenticate(s.DeletePost)).Methods("DELETE")
}

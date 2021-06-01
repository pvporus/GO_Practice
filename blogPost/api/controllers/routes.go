package controllers

func (s *Server) initializeRoutes() {

	// Home Route
	s.initializeHomeRoutes()

	// Login Route
	s.initializeLoginRoutes()

	//Users routes
	s.initializeUserRoutes()

	//Posts routes
	s.initializePostsRoutes()
}

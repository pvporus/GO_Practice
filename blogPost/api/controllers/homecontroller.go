package controllers

import (
	"net/http"

	"blogPost/api/responses"
)

//Handles the home route
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Blog Post")

}

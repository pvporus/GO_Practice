package main

func (a *App) mapRoutes() {
	//handle application routes
	a.Router.HandleFunc("/books", a.getBooks).Methods("GET")
	a.Router.HandleFunc("/book", a.addBook).Methods("POST")
	a.Router.HandleFunc("/book/{id:[0-9]+}", a.getBook).Methods("GET")
	a.Router.HandleFunc("/book/{id:[0-9]+}", a.updateBook).Methods("PUT")
	a.Router.HandleFunc("/book/{id:[0-9]+}", a.deleteBook).Methods("DELETE")
	//handlers for profiles
	a.mapProfileRoutes()
}

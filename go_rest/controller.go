package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(dbname, user, password string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.mapRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		getErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	b := book{ID: id}
	if err := b.getBook(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			getErrorResponse(w, http.StatusNotFound, "Book not found")
		default:
			getErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	getJSONResponse(w, http.StatusOK, b)
}

func getErrorResponse(w http.ResponseWriter, code int, message string) {
	getJSONResponse(w, code, map[string]string{"error": message})
}

func getJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getBooks(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getBooks(a.DB, start, count)
	if err != nil {
		getErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	getJSONResponse(w, http.StatusOK, products)
}

func (a *App) addBook(w http.ResponseWriter, r *http.Request) {
	var b book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		getErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := b.addBook(a.DB); err != nil {
		getErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	getJSONResponse(w, http.StatusCreated, b)
}

func (a *App) updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		getErrorResponse(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var b book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		getErrorResponse(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	b.ID = id

	if err := b.updateBook(a.DB); err != nil {
		getErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	getJSONResponse(w, http.StatusOK, b)
}

func (a *App) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		getErrorResponse(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	b := book{ID: id}
	if err := b.deleteBook(a.DB); err != nil {
		getErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	getJSONResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) mapRoutes() {
	a.Router.HandleFunc("/books", a.getBooks).Methods("GET")
	a.Router.HandleFunc("/book", a.addBook).Methods("POST")
	a.Router.HandleFunc("/book/{id:[0-9]+}", a.getBook).Methods("GET")
	a.Router.HandleFunc("/book/{id:[0-9]+}", a.updateBook).Methods("PUT")
	a.Router.HandleFunc("/book/{id:[0-9]+}", a.deleteBook).Methods("DELETE")
}

//Profiling, Bench mark, test coverage, load test commands
//go test -bench=. -benchmem
//go test -bench -run=XXX -cpuprofile cpu.prof .
//go tool pprof cpu.prof
//go get golang.org/x/tools/cmd/cover -- for test coverage
//go test -coverprofile testtcoverage.html fmt - get coverage
//Apache http load test cmd: ab -k -c 10 -n 100000 "http://localhost:8088/books"
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

//Create table query
const tableCreationQuery = `CREATE TABLE IF NOT EXISTS books
(
    id SERIAL,
    title TEXT NOT NULL,
	category TEXT NOT NULL,
	author TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT books_pkey PRIMARY KEY (id)
)`

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("booksDB", "postgres", "root")

	isTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

//Check whether the table exists or not
func isTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

//Clears the records in table
func clearTable() {
	a.DB.Exec("DELETE FROM books")
	a.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}

//Check empty table
func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req)

	getResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

//execute the http request
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	r := httptest.NewRecorder()
	a.Router.ServeHTTP(r, req)

	return r
}

//get the response code
func getResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, but got %d\n", expected, actual)
	}
}

//Check there is no book in table
func TestGetNonExistBook(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/book/11", nil)
	response := executeRequest(req)

	getResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Book not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Book not found'. Got '%s'", m["error"])
	}
}

//Test the adding book to books entity
func TestAddBook(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"title":"test book", "category":"test category", "autho":"author1", "price": 200.50}`)
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	getResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "test book" {
		t.Errorf("Expected book title to be 'test book'. Got '%v'", m["title"])
	}

	if m["price"] != 200.50 {
		t.Errorf("Expected product price to be '200.50'. Got '%v'", m["price"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected book ID to be '1'. Got '%v'", m["id"])
	}
}

//get the book
func TestGetBook(t *testing.T) {
	clearTable()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)

	getResponseCode(t, http.StatusOK, response.Code)
}

//add books helper function
func addBooks(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO books(title,category,author, price) VALUES($1, $2, $3, $4)", "book "+strconv.Itoa(i), "category "+strconv.Itoa(i), "author "+strconv.Itoa(i), (i+1.0)*50)
	}
}

//Check the books table update
func TestUpdateBook(t *testing.T) {

	clearTable()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	var originalBook map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBook)

	var jsonStr = []byte(`{"title":"test book updated title", "category":"test category updated", "author":"test author updated ", "price": 300.70}`)
	req, _ = http.NewRequest("PUT", "/book/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)

	getResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalBook["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalBook["id"], m["id"])
	}

	if m["name"] == originalBook["title"] {
		t.Errorf("Expected the title to be updated from '%v' to '%v'. Got '%v'", originalBook["title"], m["title"], m["name"])
	}

}

//Check the book deletion
func TestDeleteBook(t *testing.T) {
	clearTable()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	getResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/book/1", nil)
	response = executeRequest(req)

	getResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/book/1", nil)
	response = executeRequest(req)
	getResponseCode(t, http.StatusNotFound, response.Code)
}

//Bench mark the get all the books
func BenchmarkGetBooks(b *testing.B) {
	r := httptest.NewRequest("GET", "/books", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		a.getBooks(w, r)
	}
}

//Bench mark the get single book
func BenchmarkGetBookById(b *testing.B) {
	r, _ := http.NewRequest("GET", "/book/1", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		a.getBook(w, r)
	}
}

//Bench mark adding the book
func BenchmarkAddBook(b *testing.B) {
	var jsonStr = []byte(`{"title":"test book", "category":"test category", "autho":"author1", "price": 200.50}`)
	r, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		a.getBooks(w, r)
	}
}

//Bench mark updating the book
func BenchmarkUpdateBook(b *testing.B) {
	var jsonStr = []byte(`{"title":"test book updated title", "category":"test category updated", "author":"test author updated ", "price": 300.70}`)
	r, _ := http.NewRequest("PUT", "/book/1", bytes.NewBuffer(jsonStr))
	r.Header.Set("Content-Type", "application/json")
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		a.updateBook(w, r)
	}
}

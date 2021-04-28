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

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("booksDB", "postgres", "root")

	isTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func isTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM books")
	a.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS books
(
    id SERIAL,
    title TEXT NOT NULL,
	category TEXT NOT NULL,
	author TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT books_pkey PRIMARY KEY (id)
)`

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req)

	getkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	r := httptest.NewRecorder()
	a.Router.ServeHTTP(r, req)

	return r
}

func getkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, but got %d\n", expected, actual)
	}
}

func TestGetNonExistBook(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/book/11", nil)
	response := executeRequest(req)

	getkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Book not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Book not found'. Got '%s'", m["error"])
	}
}

func TestAddBook(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"title":"test book", "category":"test category", "autho":"author1", "price": 200.50}`)
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	getkResponseCode(t, http.StatusCreated, response.Code)

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

func TestGetBook(t *testing.T) {
	clearTable()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)

	getkResponseCode(t, http.StatusOK, response.Code)
}

func addBooks(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO books(title,category,author, price) VALUES($1, $2, $3, $4)", "book "+strconv.Itoa(i), "category "+strconv.Itoa(i), "author "+strconv.Itoa(i), (i+1.0)*50)
	}
}

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

	getkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalBook["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalBook["id"], m["id"])
	}

	if m["name"] == originalBook["title"] {
		t.Errorf("Expected the title to be updated from '%v' to '%v'. Got '%v'", originalBook["title"], m["title"], m["name"])
	}

}

func TestDeleteBook(t *testing.T) {
	clearTable()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	getkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/book/1", nil)
	response = executeRequest(req)

	getkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/book/1", nil)
	response = executeRequest(req)
	getkResponseCode(t, http.StatusNotFound, response.Code)
}

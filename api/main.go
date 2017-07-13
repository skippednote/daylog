package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/pat"
)

var db *sql.DB

func init() {
	db = connectDB()
	generateData()
}

func main() {

	router := pat.New()
	router.Put("/goals/{id}/entries/{eid}", updateEntry)
	router.Delete("/goals/{id}/entries/{eid}", deleteEntry)
	router.Get("/goals/{id}/entries/{eid}", entryOne)
	router.Post("/goals/{id}/entries", addEntry)
	router.Get("/goals/{id}/entries", entries)
	router.Put("/goals/{id}", updateGoal)
	router.Delete("/goals/{id}", deleteGoal)
	router.Get("/goals/{id}", goalOne)
	router.Post("/goals", addGoal)
	router.Get("/goals", goals)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

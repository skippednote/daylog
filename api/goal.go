package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func goals(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM goal;")
	if err != nil {
		log.Fatal(err)
	}
	var gs []goal
	for rows.Next() {
		var g goal
		rows.Scan(&g.ID, &g.Name, &g.Complete)
		gs = append(gs, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gs)
}

func goalOne(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	row := db.QueryRow("SELECT * FROM goal where id=$1;", id)

	var g goal
	err := row.Scan(&g.ID, &g.Name, &g.Complete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func addGoal(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var g goal
	json.NewDecoder(r.Body).Decode(&g)

	q := fmt.Sprintf("INSERT INTO goal (name) VALUES ('%s') RETURNING id", g.Name)
	var id string
	err := db.QueryRow(q).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func deleteGoal(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	q := fmt.Sprintf("DELETE FROM goal WHERE id=%s;", id)
	_, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func updateGoal(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	defer r.Body.Close()
	var g goal
	json.NewDecoder(r.Body).Decode(&g)

	_, err := db.Exec("UPDATE goal SET name=$1, complete=$2 where id=$3;", &g.Name, &g.Complete, id)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

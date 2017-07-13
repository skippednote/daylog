package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func entries(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	rows, err := db.Query("SELECT entry.id, entry.description, entry.complete, entry.created_at, entry.goal_id FROM goal INNER JOIN entry ON goal.id = entry.goal_id WHERE goal.id=$1;", id)
	if err != nil {
		log.Fatal(err)
	}
	var es []entry
	for rows.Next() {
		var e entry
		rows.Scan(&e.ID, &e.Description, &e.Complete, &e.CreatedAt, &e.GoalID)
		es = append(es, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(es)
}

func entryOne(w http.ResponseWriter, r *http.Request) {
	gid := r.URL.Query().Get(":id")
	eid := r.URL.Query().Get(":eid")
	row := db.QueryRow("SELECT entry.id, entry.description, entry.complete, entry.created_at, entry.goal_id FROM goal INNER JOIN entry ON goal.id = entry.goal_id WHERE goal.id=$1 AND entry.id=$2;", gid, eid)
	var e entry
	err := row.Scan(&e.ID, &e.Description, &e.Complete, &e.CreatedAt, &e.GoalID)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func addEntry(w http.ResponseWriter, r *http.Request) {
	gid := r.URL.Query().Get(":id")
	defer r.Body.Close()
	var e entry
	json.NewDecoder(r.Body).Decode(&e)

	q := fmt.Sprintf("INSERT INTO entry (description, goal_id) VALUES ('%s', '%s') RETURNING id", e.Description, gid)
	var id string
	err := db.QueryRow(q).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func deleteEntry(w http.ResponseWriter, r *http.Request) {
	eid := r.URL.Query().Get(":eid")

	q := fmt.Sprintf("DELETE FROM entry WHERE id=('%s');", eid)
	_, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func updateEntry(w http.ResponseWriter, r *http.Request) {
	eid := r.URL.Query().Get(":eid")
	defer r.Body.Close()
	var e entry
	json.NewDecoder(r.Body).Decode(&e)

	_, err := db.Exec("UPDATE entry SET description=$1, complete=$2 where id=$3;", &e.Description, &e.Complete, eid)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

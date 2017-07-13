package main

import (
	"fmt"
	"log"

	"github.com/icrowley/fake"
)

func generateData() {
	for i := 0; i < 10; i++ {
		g := fmt.Sprintf("INSERT INTO goal (name) VALUES ('%s') RETURNING id", fake.Word())
		var id int
		err := db.QueryRow(g).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		for j := 0; j < 10; j++ {
			e := fmt.Sprintf("INSERT INTO entry (description, goal_id) VALUES ('%s', %d);", fake.Sentence(), id)
			_, err = db.Exec(e)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

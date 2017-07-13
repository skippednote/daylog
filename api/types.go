package main

import "time"

type goal struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

type entry struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Complete    bool      `json:"complete"`
	CreatedAt   time.Time `json:"created_at"`
	GoalID      int       `json:"goal_id"`
}

package main

import (
	"encoding/json"
	"os"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func Load(filename string) ([]Todo, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}
	var todos []Todo
	if err := json.Unmarshal(file, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func Save(filename string, todos []Todo) error {
	tempFile := filename + ".tmp"
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return err
	}
	return os.Rename(tempFile, filename)
}

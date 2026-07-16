package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Todo represents a single task persisted in the JSON storage file.
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// Load reads and unmarshals the task list from filename. A missing file is
// not an error: it is treated as an empty task list, since that is the
// expected state on first run.
func Load(filename string) ([]Todo, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, fmt.Errorf("read %s: %w", filename, err)
	}
	var todos []Todo
	if err := json.Unmarshal(file, &todos); err != nil {
		return nil, fmt.Errorf("unmarshal %s: %w", filename, err)
	}
	return todos, nil
}

// Save marshals todos as JSON and writes them to filename atomically: the
// data is first written to a temporary file in the same directory, then
// moved into place with os.Rename. Since rename is atomic on the same file
// system, readers always observe either the previous complete file or the
// new one, never a partially written one.
func Save(filename string, todos []Todo) error {
	tempFile := filename + ".tmp"
	data, err := json.Marshal(todos)
	if err != nil {
		return fmt.Errorf("marshal todos: %w", err)
	}
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", tempFile, err)
	}
	if err := os.Rename(tempFile, filename); err != nil {
		return fmt.Errorf("rename %s to %s: %w", tempFile, filename, err)
	}
	return nil
}

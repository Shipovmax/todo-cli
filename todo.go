package main

import (
	"errors"
	"time"
)

// ErrTodoNotFound is returned by Complete and Delete when no task with the
// requested ID exists in the list.
var ErrTodoNotFound = errors.New("task not found")

// nextID returns the smallest ID greater than every existing task ID.
// max(existing)+1 is used instead of len(todos)+1 so that IDs stay unique
// after deletions: len alone can reissue an ID that still exists after a
// task in the middle of the slice was removed.
func nextID(todos []Todo) int {
	maxID := 0
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID + 1
}

// Add appends a new task with the given title to todos and returns the
// updated slice along with the created Todo.
func Add(todos []Todo, title string) ([]Todo, Todo) {
	newTodo := Todo{
		ID:        nextID(todos),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	return append(todos, newTodo), newTodo
}

// Complete marks the task with the given id as done and returns the updated
// slice along with the completed Todo. It returns ErrTodoNotFound if no task
// with that id exists.
func Complete(todos []Todo, id int) ([]Todo, Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			return todos, todos[i], nil
		}
	}
	return nil, Todo{}, ErrTodoNotFound
}

// Delete removes the task with the given id from todos and returns the
// updated slice along with the removed Todo. It returns ErrTodoNotFound if
// no task with that id exists.
func Delete(todos []Todo, id int) ([]Todo, Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return todos, todo, nil
		}
	}
	return nil, Todo{}, ErrTodoNotFound
}

package main

import (
	"errors"
	"time"
)

func nextID(todos []Todo) int {
	maxID := 0
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID + 1
}

func Add(todos []Todo, title string) ([]Todo, Todo) {
	newTodo := Todo{
		ID:        nextID(todos),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	return append(todos, newTodo), newTodo
}

func Complete(todos []Todo, id int) ([]Todo, Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = true
			return todos, todos[i], nil
		}
	}
	return nil, Todo{}, errors.New("задача не найдена")
}

func Delete(todos []Todo, id int) ([]Todo, Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return todos, todo, nil
		}
	}
	return nil, Todo{}, errors.New("задача не найдена")
}

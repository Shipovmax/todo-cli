// Command todo is a CLI task manager backed by a JSON file (todos.json) in
// the working directory. Supported subcommands: add, list, done, delete.
package main

import (
	"fmt"
	"os"
	"strconv"
)

const storageFile = "todos.json"

func main() {
	todos, err := Load(storageFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %s\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "error: please specify a command: add, list, done, delete\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "error: the add command requires task text\n")
			os.Exit(1)
		}
		todos, newTodo := Add(todos, os.Args[2])
		if err := Save(storageFile, todos); err != nil {
			fmt.Fprintf(os.Stderr, "error saving: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Added: [%d] %s\n", newTodo.ID, newTodo.Title)
	case "list":
		for _, todo := range todos {
			status := "[ ]"
			if todo.Done {
				status = "[x]"
			}
			fmt.Printf("%s  %-5d %s\n", status, todo.ID, todo.Title)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "error: please specify a valid task ID\n")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprint(os.Stderr, "error: ID must be a number\n")
			os.Exit(1)
		}
		todos, completed, err := Complete(todos, id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		if err := Save(storageFile, todos); err != nil {
			fmt.Fprintf(os.Stderr, "error saving: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Completed: [%d] %s\n", completed.ID, completed.Title)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "error: please specify a valid task ID\n")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprint(os.Stderr, "error: ID must be a number\n")
			os.Exit(1)
		}
		todos, deleted, err := Delete(todos, id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		if err := Save(storageFile, todos); err != nil {
			fmt.Fprintf(os.Stderr, "error saving: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted: [%d] %s\n", deleted.ID, deleted.Title)
	default:
		fmt.Fprintf(os.Stderr, "error: unknown command %s\n", os.Args[1])
		os.Exit(1)
	}
}

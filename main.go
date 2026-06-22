package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	todos, err := Load("todos.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка чтения файла: %s\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "ошибка: укажите команду: add, list, done, delete\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "ошибка: укажите текст задачи\n")
			os.Exit(1)
		}
		todos, newTodo := Add(todos, os.Args[2])
		if err := Save("todos.json", todos); err != nil {
			fmt.Fprintf(os.Stderr, "ошибка сохранения: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Добавлено: [%d] %s\n", newTodo.ID, newTodo.Title)
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
			fmt.Fprintf(os.Stderr, "ошибка: укажите правильный номер задачи\n")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprint(os.Stderr, "ошибка: ID должен быть числом\n")
			os.Exit(1)
		}
		todos, completed, err := Complete(todos, id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ошибка: %s\n", err)
			os.Exit(1)
		}
		if err := Save("todos.json", todos); err != nil {
			fmt.Fprintf(os.Stderr, "ошибка сохранения: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Выполнено: [%d] %s\n", completed.ID, completed.Title)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "ошибка: укажите правильный номер задачи\n")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprint(os.Stderr, "ошибка: ID должен быть числом\n")
			os.Exit(1)
		}
		todos, deleted, err := Delete(todos, id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ошибка: %s\n", err)
			os.Exit(1)
		}
		if err := Save("todos.json", todos); err != nil {
			fmt.Fprintf(os.Stderr, "ошибка сохранения: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Удалено: [%d] %s\n", deleted.ID, deleted.Title)
	default:
		fmt.Fprintf(os.Stderr, "ошибка: неизвестная команда %s\n", os.Args[1])
		os.Exit(1)
	}
}

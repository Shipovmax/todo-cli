package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	todos, _ := Load("todos.json")

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
			fmt.Printf("%s %d %s\n", status, todo.ID, todo.Title)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "ошибка: укажите правильный номер задачи\n")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprint(os.Stderr, "ошибка: ID Должен быть числом\n")
			os.Exit(1)
		}
		todos, err = Complete(todos, id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ошибка: %s\n", err)
			os.Exit(1)
		}
		if err := Save("todos.json", todos); err != nil {
			fmt.Fprintf(os.Stderr, "ошибка сохранения: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Выполнено: [%d]\n", id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "ошибка: укажите правильный номер задачи\n")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprint(os.Stderr, "ошибка: ID Должен быть числом\n")
			os.Exit(1)
		}
		todos, err = Delete(todos, id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ошибка: %s\n", err)
			os.Exit(1)
		}
		if err := Save("todos.json", todos); err != nil {
			fmt.Fprintf(os.Stderr, "ошибка сохранения: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Удалено: [%d]\n", id)
	default:
		fmt.Fprintf(os.Stderr, "ошибка: неизвестная команда %s\n", os.Args[1])
		os.Exit(1)
	}
}

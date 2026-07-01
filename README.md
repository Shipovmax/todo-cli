# todo-cli — A Persistent TODO Manager

> A CLI application for task management with JSON file storage. Learning Project #3 as part of preparation for a Go Backend Developer role.

---

## For the Recruiter

### What it is and Why

The third project in the roadmap marks the transition from stateless computations to a stateful application. Tasks are persisted across restarts using a JSON file: this represents a minimal model of how any backend with persistent storage operates, using the file system instead of PostgreSQL.

The primary goal is to master file I/O in Go, struct serialization using `encoding/json`, slice operations (appending, filtering, searching by ID), and proper state management via a single data file. All of these are direct analogs of database operations: read → modify → write.

The project demonstrates a solid understanding of layer separation: CLI argument parsing is decoupled from business logic, which is further separated from the storage layer. This exact three-tier architecture is what scales into real-world production services.

### What This Project Demonstrates

| Skill | Implementation |
|---|---|
| File I/O | `os.ReadFile`, `os.WriteFile`, creating the file if it does not exist |
| JSON Serialization | `encoding/json`, struct tags, `json.Marshal` / `json.Unmarshal` |
| Working with Structs | `Todo` struct with fields: `ID`, `Title`, `Done`, `CreatedAt` |
| Slice Operations | appending, filtering by condition, searching by ID |
| CLI with Subcommands | `os.Args` parsing for subcommands: `add` / `list` / `done` / `delete` |
| Layer Separation | `storage.go` (I/O) + `todo.go` (business logic) + `main.go` (CLI) |
| Formatted Output | `fmt.Printf` with padding/alignment for the task table |

### Stack

- **Language:** Go 1.22+
- **Dependencies:** Standard library only
- **Storage:** `todos.json` in the working directory
- **Platform:** Linux / macOS / Windows

---

## For the Developer

### Architectural Decisions

#### Why a JSON file instead of SQLite or in-memory?

A JSON file provides minimal persistence with zero external dependencies. The goal of this project is to understand the read → modify → write pattern, which is identical for both files and databases. SQLite would introduce a database driver and SQL, distracting from the core objective. An in-memory store wouldn't persist data between runs, which defeats the educational value.

#### Why `os.ReadFile` / `os.WriteFile` instead of `os.Open` + buffered reading?

For a file that is only a few kilobytes in size (a task list), buffering is unnecessary. `ReadFile` and `WriteFile` are idiomatic Go for handling small files in their entirety. `bufio` will come in handy in subsequent projects when reading large files line-by-line.

#### Why subcommands via `os.Args` instead of flags via `flag`?

`add`, `list`, `done`, and `delete` are commands, not flags. Semantically, they are closer to subcommands (like `git commit`, `docker run`) rather than flags (like `-v`, `--output`). Utilizing `os.Args[1]` as a subcommand is the correct model for CLI tools.

#### Why is the ID generated as `max(existing IDs) + 1` instead of `len(todos) + 1`?

Using `len` causes collisions after a task is deleted. For example: if you add 3 tasks, delete the second one, and then add a new one, `len` will yield ID=3, which already exists. Using `max+1` guarantees uniqueness.

#### Why atomic write: first to a temp file, then rename?

```go
// Unsafe: if the process crashes during writing, the file gets corrupted
os.WriteFile("todos.json", data, 0644)

// Safe: rename is atomic at the OS level
os.WriteFile("todos.json.tmp", data, 0644)
os.Rename("todos.json.tmp", "todos.json")

```

An `os.Rename` operation on the same file system is atomic. You either get the old file or the new one, completely eliminating any intermediate corrupted state.

### Structure

```
todo-cli/
├── main.go       # CLI: parsing os.Args, invoking commands, printing results
├── todo.go       # Business logic: Add, List, Complete, Delete, nextID
├── storage.go    # I/O layer: Load (read+unmarshal), Save (marshal+write)
├── go.mod
└── README.md

```

### Installation and Setup

```bash
git clone [https://github.com/Shipovmax/todo-cli](https://github.com/Shipovmax/todo-cli)
cd todo-cli
go build -o todo .

```

### Usage

```bash
./todo add <task text>     # add a task
./todo list                   # list all tasks
./todo done <id>              # mark a task as completed
./todo delete <id>            # delete a task

```

### Examples

```bash
./todo add "Learn encoding/json"
# Added: [1] Learn encoding/json

./todo add "Write storage.go"
# Added: [2] Write storage.go

./todo list
# ID  Status  Task
# 1   [ ]     Learn encoding/json
# 2   [ ]     Write storage.go

./todo done 1
# Completed: [1] Learn encoding/json

./todo list
# ID  Status  Task
# 1   [x]     Learn encoding/json
# 2   [ ]     Write storage.go

./todo delete 2
# Deleted: [2] Write storage.go

```

### Error Handling

```bash
./todo done 99
# stderr: error: task with ID 99 not found
# exit code: 1

./todo delete 99
# stderr: error: task with ID 99 not found
# exit code: 1

./todo
# stderr: error: please specify a command: add / list / done / delete
# exit code: 1

./todo add
# stderr: error: the add command requires task text
# exit code: 1

./todo done abc
# stderr: error: ID must be a number
# exit code: 1

```

### Running Without Building

```bash
go run . add "First task"
go run . list

```



Жду следующий файл!

```

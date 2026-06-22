# Task #3 — TODO CLI с персистентностью

## Цель

Написать CLI-приложение для управления задачами, которое сохраняет данные в JSON-файл между запусками. Главная учебная цель — освоить файловый I/O (`os.ReadFile`/`os.WriteFile`), сериализацию структур через `encoding/json`, и операции над слайсами структур (добавление, поиск, фильтрация). Это прямая аналогия паттерна read → modify → write, который используется при работе с любым хранилищем данных.

---

## Acceptance Criteria

- [ ] `./todo add "купить молоко"` → stdout: `Добавлено: [1] купить молоко`
- [ ] `./todo list` → таблица с ID, статусом `[ ]`/`[x]` и текстом задачи
- [ ] `./todo done 1` → stdout: `Выполнено: [1] купить молоко`, задача помечена `[x]` в list
- [ ] `./todo delete 1` → stdout: `Удалено: [1] купить молоко`, задача исчезает из list
- [ ] После перезапуска программы данные сохранены в `todos.json`
- [ ] `./todo done 99` → stderr: `ошибка: задача с ID 99 не найдена`, exit code 1
- [ ] `./todo` (без аргументов) → stderr: ошибка с подсказкой команд, exit code 1
- [ ] `./todo add` (без текста) → stderr: ошибка, exit code 1
- [ ] `./todo done abc` → stderr: `ошибка: ID должен быть числом`, exit code 1
- [ ] Удаление задачи не ломает ID уже существующих задач
- [ ] `go vet ./...` проходит без предупреждений
- [ ] `go.mod` содержит только `module` и `go` директивы

---

## Технические требования

### Обязательно

| Требование | Детали |
|---|---|
| Хранилище | файл `todos.json` в рабочей директории |
| Чтение файла | `os.ReadFile` + `json.Unmarshal` |
| Запись файла | `json.Marshal` + атомарная запись: temp-файл + `os.Rename` |
| Структура задачи | `Todo{ID int, Title string, Done bool, CreatedAt time.Time}` |
| JSON теги | `json:"id"`, `json:"title"`, `json:"done"`, `json:"created_at"` |
| ID генерация | `max(existing IDs) + 1`, не `len + 1` |
| Субкоманды | `os.Args[1]` как subcommand: `add`, `list`, `done`, `delete` |
| Парсинг ID | `strconv.Atoi` с обработкой ошибки |
| Разбивка файлов | минимум 3 файла: `main.go`, `todo.go`, `storage.go` |
| Вывод ошибок | `fmt.Fprintf(os.Stderr, ...)` + `os.Exit(1)` |

### Запрещено

- `panic` для обработки ошибок — только `error` return
- Сторонние библиотеки — только стандартная библиотека
- `len(todos) + 1` для генерации ID — приводит к коллизиям после удаления
- Прямая запись в `todos.json` без temp-файла + rename — небезопасно при сбое

---

## Темы Go, которые ты прокачиваешь

> Это не просто список — это checklist того, что **обязан использовать** в реализации.

- **`os.ReadFile` / `os.WriteFile`** — читать весь файл в `[]byte` и записывать обратно
- **`os.Rename`** — атомарная замена файла через temp + rename
- **`encoding/json`** — `json.Marshal(todos)` для сериализации, `json.Unmarshal(data, &todos)` для десериализации
- **struct tags** — `json:"field_name"` управляют именами ключей в JSON
- **`time.Time`** — поле `CreatedAt` с тегом `json:"created_at"`, автоматическая сериализация в ISO 8601
- **операции над `[]Todo`** — append для добавления, цикл + условие для поиска/фильтрации
- **`strconv.Atoi`** — парсинг строкового ID в int с обработкой ошибки
- **`os.IsNotExist`** — проверка что файл ещё не создан при первом запуске
- **`fmt.Printf` с форматированием** — выравнивание колонок в выводе `list`

---

## Структура файлов

```
todo-cli/
├── main.go       # парсинг os.Args, switch по subcommand, вывод результата
├── todo.go       # Todo struct + функции: Add, List, Complete, Delete, nextID
├── storage.go    # Load(filename) ([]Todo, error), Save(filename, []Todo) error
├── go.mod        # module github.com/Shipovmax/todo-cli
└── README.md
```

---

## Подсказки по архитектуре

```go
// storage.go — только I/O, ничего про бизнес-логику
func Load(filename string) ([]Todo, error)
func Save(filename string, todos []Todo) error

// todo.go — бизнес-логика, не знает про файлы и CLI
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
}

func Add(todos []Todo, title string) ([]Todo, Todo)
func Complete(todos []Todo, id int) ([]Todo, error)
func Delete(todos []Todo, id int) ([]Todo, error)
func nextID(todos []Todo) int

// main.go — только CLI, делегирует всё остальное
func main() {
    todos, _ := Load("todos.json")
    switch os.Args[1] {
    case "add": ...
    case "list": ...
    case "done": ...
    case "delete": ...
    }
    Save("todos.json", todos)
}
```

> Обрати внимание: `Load` и `Save` вызываются в `main.go`, а не внутри бизнес-функций. Бизнес-логика работает только с `[]Todo` в памяти — это делает её тестируемой без файловой системы.

---

## Definition of Done

1. Все Acceptance Criteria выполнены
2. Код запушен на GitHub в репозиторий `todo-cli`
3. README.md в репозитории соответствует шаблону проекта
4. Ты можешь объяснить каждую строку кода вслух без подглядывания

---

## Следующий шаг после сдачи

После ревью переходим к **Task #4 — TODO HTTP API**: та же бизнес-логика из этого проекта, но exposed через REST API — stateless хендлеры, in-memory хранилище с `sync.RWMutex`, полный CRUD через HTTP.

# Go ToDo CLI Using Cobra

A simple yet powerful terminal-based ToDo List application written in Golang using the Cobra CLI framework and MySQL for persistent task storage. The app supports adding, completing, deleting, updating, and listing tasks — all from the command line.

---

## Features

- Add new tasks with titles and descriptions
- List all tasks in a clean tabular format
- Update task details (title or description)
- Delete tasks
- Mark tasks as completed
- See human-readable timestamps using `timediff`
- MySQL database integration using GORM ORM

---

## Tech Stack

- **Language:** Go (Golang)
- **Framework:** Cobra CLI
- **Database:** MySQL
- **ORM:** GORM
- **Time Formatter:** [`timediff`](https://github.com/mergestat/timediff)
- **Output Formatter:** `tabwriter`

---

## Project Structure

```

goToDo/
│
├── cmd/                 # Cobra commands (add, list, update, delete, complete)
│   └── ...
│
├── config/              # config.yaml and config loader
│   └── config.go
│
├── internal/            # Business logic
│   ├── db/              # DB connection code
│   ├── models/          # Task model
│   └── utils/           # Helper functions
│
├── docs/                # Documentation folder
│   └── setup.md         # Setup guide
│
├── go.mod / go.sum
└── main.go

````

---

## Setup Instructions

Full setup guide available at:  
[`docs/setup.md`](docs/setup.md)

---

## Usage

```bash
# Add a new task
todo add -t "Buy groceries" -d "Milk, eggs, bread"

# List all tasks
todo list

# Update a task
todo update -i 1 -t "Buy snacks" -d "Chips, soda"

# Mark task as completed
todo markcompleted -i 1

# Delete a task
todo delete -i 1
````

---

## What I Learned

> Building this project helped me gain confidence in using Go’s Cobra CLI, structuring terminal-based projects, handling SQL with GORM, managing timestamps manually, and improving code clarity. I now understand how important logic-building is and plan to improve on it moving forward.

---

## License

MIT License — Feel free to use and modify.

---

## Acknowledgements

Thanks to [timediff](https://github.com/mergestat/timediff) for human-friendly timestamps, and the Go + GORM + Cobra communities for solid documentation.
---
Made by Ranit Santra

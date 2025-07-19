# 🛠️ Project Setup Guide

Welcome to the **Go ToDo CLI App** setup guide. This document will walk you through all the steps required to run the project on your local machine.

---

## 📦 Prerequisites

Ensure you have the following tools installed before starting:

- [Go (v1.20+)](https://golang.org/dl/)
- [MySQL Server](https://dev.mysql.com/downloads/mysql/)
- Git

---

## 🧾 MySQL Database Setup

1. **Start your MySQL server.**

2. **Create a database:**

```sql
CREATE DATABASE go_todo_db;
````

---

## ⚙️ Project Configuration

1. Go to the `config/config.yaml` file and update your DB credentials:

```yaml
db:
  host: "localhost"
  port: "3306"
  user: "root"            # or "todo_user"
  password: "your_password"
  name: "go_todo_db"
```

> ✅ Ensure the `config.yaml` file is excluded from version control via `.gitignore`.

---

## 📂 Project Structure

```text
goToDo/
│
├── cmd/                   # Cobra command definitions (add, list, update, delete, etc.)
├── config/                # Contains config.yaml and config loader
├── internal/
│   ├── db/                # Database connection and auto-migration
│   ├── models/            # Task struct and DB model
│   ├── utils/             # Time formatting, helpers
├── docs/                  # Markdown docs like this one
├── go.mod
├── main.go
└── README.md
```

---

## 🚀 Run the App

### Step 1: Download Dependencies

```bash
go mod tidy
```

### Step 2: Build the Binary

```bash
go build -o todo
```

This will generate a binary called `todo`.

### Step 3: Use the CLI

#### Add a Task

```bash
./todo add -t "Buy Groceries" -d "Milk, Bread, Eggs"
```

#### List All Tasks

```bash
./todo list
```

#### Update a Task

```bash
./todo update -i 1 -t "Buy Groceries (Updated)"
```

#### Mark as Completed

```bash
./todo markcompleted -i 1
```

#### Delete a Task

```bash
./todo delete -i 1
```

---

## 🧪 Sample Output

```text
ID  Title           Completed  Created At         Updated At         Completed At
1   Buy Groceries   ✅          2 minutes ago      1 minute ago       1 minute ago
```

---

## 🧹 Clean & Rebuild

If you need to clean and rebuild:

```bash
go clean
go build -o todo
```

---

## 📌 Notes

* Tasks are stored persistently in the MySQL database.
* Task times are recorded in `UTC` unless you modify the time zone settings.
* Make sure the MySQL server is running when using the CLI.

---
Happy coding! 🚀

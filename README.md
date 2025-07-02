# api-crud-assessment

# Task Management API – Golang (Gin + PostgreSQL)

A clean-architecture RESTful API in Go using the **Gin framework** and **PostgreSQL**, built for the API Developer Skill Test. This API manages "Tasks" with fields for `name`, `description`, and `priority`.

---

## 🧱 Tech Stack

- Go 1.20+
- Gin Web Framework
- PostgreSQL (Supabase-hosted)
- `database/sql` + `lib/pq`
- Dependency Injection (no interface layer between service/handler)
- .env-based configuration

---

## 📁 Project Structure

go-tasks-api/
├── config/ # DB setup
├── internal/
│ ├── handler/ # HTTP handlers
│ ├── model/ # Task struct
│ ├── repository/
| | ├── repo_postgres/ # DB layer
│ ├── router/ # Routing definitions
│ └── usecase/ # Business logic
├── schema.sql # PostgreSQL table definition
├── .env # Environment config
├── go.mod
└── README.md

## DB Schema

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR NULL,
    description VARCHAR NULL,
    priority INT NULL
);
## API Endpoints

| Method | Endpoint    | Description     |
| ------ | ----------- | --------------- |
| GET    | /tasks/list | List all tasks  |
| GET    | /tasks/:id  | Get task by ID  |
| POST   | /tasks      | Create new task |
| PUT    | /tasks/:id  | Update task     |
| DELETE | /tasks/:id  | Delete task     |


Tested with Postman. Screenshots attached in the /screenshots folder.
Added the postman collection JSON as well

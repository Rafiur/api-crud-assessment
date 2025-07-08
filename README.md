# api-crud-assessment

# Task Management API – Golang (Gin + PostgreSQL)

A clean-architecture RESTful API built in **Go (Gin framework)** with a **hosted PostgreSQL database**. This project was developed for the API Developer Skill Test and manages a `tasks` resource with fields:

- `id` (integer, primary key)
- `name` (varchar)
- `description` (varchar)
- `priority` (integer)

---

## 🚀 How to Run This Project

### 🧱 Requirements

- Go 1.20+
- Internet connection (uses an online hosted PostgreSQL database)
- `git` and `Postman` (for API testing)

---

### 🔧 1. Clone the Repository

```bash
git clone https://github.com/yourusername/api-crud-assessment.git
cd api-crud-assessment
```
### 🔧 2. Install Go Dependencies
```bash 
go mod tidy
```
3. Set up .env File
Create a .env file in the root directory with the following content:

```env
DATABASE_URL=postgresql://neondb_owner:npg_hC84wjkTtXJB@ep-muddy-star-a8e1lz94-pooler.eastus2.azure.neon.tech/neondb?sslmode=require&channel_binding=require
```
✅ This connects to an online database hosted via Neon.tech.

4. Run the Server
```bash
go run cmd/main.go
```
By default, the server runs at:
```bash
http://localhost:3000
```

Just in Case: Example Local Connection (Optional)
If you prefer to test with a local PostgreSQL instance, use this example:
```env
DATABASE_URL=postgres://postgres:yourpassword@localhost:5432/tasksdb?sslmode=disable
```
You can create the table using the provided sql:

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR NULL,
    description VARCHAR NULL,
    priority INT NULL
);
```


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

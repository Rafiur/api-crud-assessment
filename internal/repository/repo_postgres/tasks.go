package repo_postgres

import (
	"database/sql"
	"errors"
	"github.com/Rafiur/api-crud-assessment/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (repo *TaskRepository) Create(task *model.Task) (*model.Task, error) {
	err := repo.db.QueryRow("INSERT INTO tasks (name, description, priority) VALUES ($1, $2, $3) RETURNING id",
		task.Name, task.Description, task.Priority).Scan(&task.ID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (repo *TaskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := repo.db.QueryRow("SELECT id, name, description, priority FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Name, &task.Description, &task.Priority)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (repo *TaskRepository) List() ([]model.Task, error) {
	rows, err := repo.db.Query("SELECT id, name, description, priority FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []model.Task{}
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Priority); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo *TaskRepository) Update(task *model.Task) (*model.Task, error) {
	res, err := repo.db.Exec("UPDATE tasks SET name=$1, description=$2, priority=$3 WHERE id=$4", task.Name, task.Description, task.Priority, task.ID)
	if err != nil {
		return nil, err
	}
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		return nil, errors.New("no rows updated")
	}
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (repo *TaskRepository) Delete(id int) error {
	res, err := repo.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		return errors.New("no rows deleted")
	}
	if err != nil {
		return err
	}
	return nil
}

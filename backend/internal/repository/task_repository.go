package repository

import (
    "database/sql"
    "backend/internal/task"
)

type TaskRepository interface {
    Create(task *task.Task) error
    FindByID(id int) (*task.Task, error)
    Update(task *task.Task) error
    Delete(id int) error
    FindAll() ([]*task.Task, error)
}


type taskRepository struct {
    db *sql.DB
}


func NewTaskRepository(db *sql.DB) TaskRepository {
    return &taskRepository{db: db}
}

func (r *taskRepository) Create(t *task.Task) error {
    query := `INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)`
    _, err := r.db.Exec(query, t.Title, t.Description, t.Completed)
    return err
}

func (r *taskRepository) FindByID(id int) (*task.Task, error) {
    query := `SELECT id, title, description, completed, created_at, updated_at FROM tasks WHERE id = ?`
    row := r.db.QueryRow(query, id)
    t := &task.Task{}
    err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return t, nil
}

func (r *taskRepository) Update(t *task.Task) error {
    query := `UPDATE tasks SET title = ?, description = ?, completed = ? WHERE id = ?`
    _, err := r.db.Exec(query, t.Title, t.Description, t.Completed, t.ID)
    return err
}

func (r *taskRepository) Delete(id int) error {
    query := `DELETE FROM tasks WHERE id = ?`
    _, err := r.db.Exec(query, id)
    return err
}

func (r *taskRepository) FindAll() ([]*task.Task, error) {
    query := `SELECT id, title, description, completed, created_at, updated_at FROM tasks`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []*task.Task
    for rows.Next() {
        t := &task.Task{}
        err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, t)
    }
    return tasks, nil
}
package main

import "database/sql"

type Store interface {
	CreateUser() error
	CreateTask(task *Task) (*Task, error)
	GetTask(id string) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (storage *Storage) CreateUser() error {
	return nil
}

func (storage *Storage) CreateTask(task *Task) (*Task, error) {
	rows, err := storage.db.Exec(`INSERT INTO tasks (name, status, projectId, assignedToId) VALUES (?, ?, ?, ?)`, task.Name, task.Status, task.ProjectId, task.AssignedToId)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.Id = id
	return task, nil
}

func (storage *Storage) GetTask(id string) (*Task, error) {
	var task Task
	err := storage.db.QueryRow(`SELECT id, name, status, projectId, assignedToId FROM tasks WHERE id = ?`, id).Scan(&task.Id, &task.Name, &task.Status, &task.ProjectId, &task.AssignedToId)

	return &task, err
}

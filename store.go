package main

import "database/sql"

type Store interface {
	CreateUser(user *User) (*User, error)
	GetUserById(id string) (*User, error)
	CreateTask(task *Task) (*Task, error)
	GetTask(id string) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (storage *Storage) CreateUser(
	user *User,
) (*User, error) {
	row, err := storage.db.Exec(`INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)`, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = id
	return user, nil

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

func (storage *Storage) GetUserById(id string) (*User, error) {
	var user User
	err := storage.db.QueryRow(`SELECT id, firstName, lastName, email, password, createdAt FROM users WHERE id = ?`, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)

	return &user, err
}

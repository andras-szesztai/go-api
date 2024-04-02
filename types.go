package main

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectId    int64     `json:"projectId"`
	AssignedToId int64     `json:"assignedToId"`
	CreateAt     time.Time `json:"createAt"`
}

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

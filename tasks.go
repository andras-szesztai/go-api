package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskService struct {
	store Store
}

func NewTasksService(store Store) *TaskService {
	return &TaskService{store: store}
}

func (ts *TaskService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", ts.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", ts.handleGetTask).Methods("GET")
}

func (ts *TaskService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}
	defer r.Body.Close()

	var task *Task
	if err = json.Unmarshal(body, &task); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if createdTask, err := ts.store.CreateTask(task); err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	} else {
		WriteJson(w, http.StatusCreated, createdTask)
	}

}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errors.New("task name is required")
	}
	if task.ProjectId == 0 {
		return errors.New("task project id is required")
	}
	if task.AssignedToId == 0 {
		return errors.New("task assigned to id is required")
	}
	return nil
}

func (ts *TaskService) handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	task, err := ts.store.GetTask(id)
	if err != nil {
		WriteJson(w, http.StatusNotFound, ErrorResponse{Error: "Task not found"})
		return
	}
	WriteJson(w, http.StatusOK, task)
}

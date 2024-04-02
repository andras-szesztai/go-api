package main

type MockStore struct{}

func (ms *MockStore) CreateUser() error {
	return nil
}

func (ms *MockStore) CreateTask(task *Task) (*Task, error) {
	return &Task{}, nil
}

func (ms *MockStore) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}

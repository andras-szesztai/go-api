package main

type MockStore struct{}

func (ms *MockStore) CreateUser(
	user *User,
) (*User, error) {
	return &User{}, nil
}

func (ms *MockStore) CreateTask(task *Task) (*Task, error) {
	return &Task{}, nil
}

func (ms *MockStore) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}

func (ms *MockStore) GetUserById(id string) (*User, error) {
	return &User{}, nil
}

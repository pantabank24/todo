package todoRepository

import (
	"github.com/pantabank24/todo/models"
	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (m *TodoRepositoryMock) CreateTodo(req *models.TodoReq) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *TodoRepositoryMock) ListTodos(filter *models.TodosFilter) ([]*models.TodoRes, error) {
	args := m.Called(filter)
	return args.Get(0).([]*models.TodoRes), args.Error(1)
}

func (m *TodoRepositoryMock) UpdateTodos(req *models.TodoReq, uid string) error {
	args := m.Called(req, uid)
	return args.Error(0)
}

func (m *TodoRepositoryMock) TodoDone(uid string) error {
	args := m.Called(uid)
	return args.Error(0)
}

func (m *TodoRepositoryMock) RemoveTodo(uid string) error {
	args := m.Called(uid)
	return args.Error(0)
}

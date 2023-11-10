package unittest

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/pantabank24/todo/models"
	"github.com/pantabank24/todo/modules/todo/todoRepository"
	"github.com/pantabank24/todo/modules/todo/todoUsecase"
	"github.com/stretchr/testify/assert"
)

type (
	testCreateTodo struct {
		req         *models.TodoReq
		description string
		expected    []*models.TodoRes
		isErr       bool
	}

	testUpdateTodo struct {
		req         *models.TodoReq
		uid         string
		description string
		expected    []*models.TodoRes
		isErr       bool
	}

	testTodoDone struct {
		uid         string
		description string
		expected    []*models.TodoRes
		isErr       bool
	}

	testRemoveTodo struct {
		uid         string
		description string
		expected    []*models.TodoRes
		isErr       bool
	}
)

func TestCreateTodo(t *testing.T) {
	repoMock := new(todoRepository.TodoRepositoryMock)
	usecase := todoUsecase.NewTodoUsecase(repoMock)

	tests := []testCreateTodo{
		{
			req: &models.TodoReq{
				Title:       "test001",
				Description: "this is unit test : pass",
				Image:       "https://image001.com",
			},
			description: "test create todo : success case",
			expected: []*models.TodoRes{
				{
					ID:          "550e8400-e29b-41d4-a716-446655440000",
					Title:       "test001",
					Description: "this is unit test : pass",
					CreatedAt:   time.Time{},
					Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
					Status:      "IN_PROGRESS",
				},
			},
			isErr: false,
		},
	}

	repoMock.On("CreateTodo", &models.TodoReq{
		Title:       "test001",
		Description: "this is unit test : pass",
		Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
	}).Return(nil)

	repoMock.On("ListTodos", &models.TodosFilter{
		Order:   "created_at",
		OrderBy: "DESC",
	}).Return([]*models.TodoRes{
		{
			ID:          "550e8400-e29b-41d4-a716-446655440000",
			Title:       "test001",
			Description: "this is unit test : pass",
			CreatedAt:   time.Time{},
			Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
			Status:      "IN_PROGRESS",
		},
	}, nil)

	for i, test := range tests {
		fmt.Printf("case -> %d : %s\n", i+1, test.description)

		result, err := usecase.CreateTodo(test.req)

		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestUpdateTodo(t *testing.T) {
	repoMock := new(todoRepository.TodoRepositoryMock)
	usecase := todoUsecase.NewTodoUsecase(repoMock)

	tests := []testUpdateTodo{
		{
			req: &models.TodoReq{
				Title:       "update_test001",
				Description: "this is unit test",
				Image:       "https://image001.com",
			},
			uid:         "550e8400-e29b-41d4-a716-446655440000",
			description: "test update todo : success case",
			expected: []*models.TodoRes{
				{
					ID:          "550e8400-e29b-41d4-a716-446655440000",
					Title:       "update_test001",
					Description: "this is unit test",
					CreatedAt:   time.Time{},
					Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
					Status:      "IN_PROGRESS",
				},
			},
			isErr: false,
		},
		{
			req: &models.TodoReq{
				Title:       "update_test001",
				Description: "this is unit test",
				Image:       "https://image001.com",
			},
			uid:         "abcdefgh",
			description: "test create todo : failed case (uid is wrong)",
			expected:    nil,
			isErr:       true,
		},
	}

	repoMock.On("UpdateTodos", &models.TodoReq{
		Title:       "update_test001",
		Description: "this is unit test",
		Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
	}, "550e8400-e29b-41d4-a716-446655440000").Return(nil)

	repoMock.On("UpdateTodos", &models.TodoReq{
		Title:       "update_test001",
		Description: "this is unit test",
		Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
	}, "abcdefgh").Return(errors.New("error: id not found : abcdefgh"))

	repoMock.On("ListTodos", &models.TodosFilter{
		Order:   "created_at",
		OrderBy: "DESC",
	}).Return([]*models.TodoRes{
		{
			ID:          "550e8400-e29b-41d4-a716-446655440000",
			Title:       "update_test001",
			Description: "this is unit test",
			CreatedAt:   time.Time{},
			Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
			Status:      "IN_PROGRESS",
		},
	}, nil)

	for i, test := range tests {
		fmt.Printf("case -> %d : %s\n", i+1, test.description)

		result, err := usecase.UpdateTodo(test.req, test.uid)

		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestTodoDone(t *testing.T) {
	repoMock := new(todoRepository.TodoRepositoryMock)
	usecase := todoUsecase.NewTodoUsecase(repoMock)

	tests := []testTodoDone{
		{
			uid:         "550e8400-e29b-41d4-a716-446655440000",
			description: "test todo done : success case",
			expected: []*models.TodoRes{
				{
					ID:          "550e8400-e29b-41d4-a716-446655440000",
					Title:       "test001_completed",
					Description: "this is unit test",
					CreatedAt:   time.Time{},
					Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
					Status:      "COMPLETED",
				},
			},
			isErr: false,
		},
		{
			uid:         "abcdefgh",
			description: "test todo done : failed case (uid is wrong)",
			expected:    nil,
			isErr:       true,
		},
	}

	repoMock.On("TodoDone", "550e8400-e29b-41d4-a716-446655440000").Return(nil)

	repoMock.On("TodoDone", "abcdefgh").Return(errors.New("error: id not found : abcdefgh"))

	repoMock.On("ListTodos", &models.TodosFilter{
		Order:   "created_at",
		OrderBy: "DESC",
	}).Return([]*models.TodoRes{
		{
			ID:          "550e8400-e29b-41d4-a716-446655440000",
			Title:       "test001_completed",
			Description: "this is unit test",
			CreatedAt:   time.Time{},
			Image:       "aHR0cHM6Ly9pbWFnZTAwMS5jb20=",
			Status:      "COMPLETED",
		},
	}, nil)

	for i, test := range tests {
		fmt.Printf("case -> %d : %s\n", i+1, test.description)

		result, err := usecase.TodoDone(test.uid)

		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestRemoveTodo(t *testing.T) {
	repoMock := new(todoRepository.TodoRepositoryMock)
	usecase := todoUsecase.NewTodoUsecase(repoMock)

	tests := []testRemoveTodo{
		{
			uid:         "550e8400-e29b-41d4-a716-446655440000",
			description: "test remove todo : success case",
			expected:    []*models.TodoRes{},
			isErr:       false,
		},
		{
			uid:         "abcdefgh",
			description: "test remove todo : failed case (uid is wrong)",
			expected:    nil,
			isErr:       true,
		},
	}

	repoMock.On("RemoveTodo", "550e8400-e29b-41d4-a716-446655440000").Return(nil)

	repoMock.On("RemoveTodo", "abcdefgh").Return(errors.New("error: id not found : abcdefgh"))

	repoMock.On("ListTodos", &models.TodosFilter{
		Order:   "created_at",
		OrderBy: "DESC",
	}).Return([]*models.TodoRes{}, nil)

	for i, test := range tests {
		fmt.Printf("case -> %d : %s\n", i+1, test.description)

		result, err := usecase.RemoveTodo(test.uid)

		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			assert.Equal(t, test.expected, result)
		}
	}
}

package todoUsecase

import (
	"encoding/base64"
	"strings"

	"github.com/pantabank24/todo/models"
	"github.com/pantabank24/todo/modules/todo/todoRepository"
)

type (
	TodoUsecaseService interface {
		CreateTodo(req *models.TodoReq) ([]*models.TodoRes, error)
		ListTodos(filter *models.TodosFilter) ([]*models.TodoRes, error)
		UpdateTodo(req *models.TodoReq, uid string) ([]*models.TodoRes, error)
		TodoDone(uid string) ([]*models.TodoRes, error)
		RemoveTodo(uid string) ([]*models.TodoRes, error)
	}

	todoUsecase struct {
		todoRepository todoRepository.TodoRepositoryService
	}
)

func NewTodoUsecase(todoRepository todoRepository.TodoRepositoryService) TodoUsecaseService {
	return &todoUsecase{todoRepository}
}

func (u *todoUsecase) CreateTodo(req *models.TodoReq) ([]*models.TodoRes, error) {
	textBytes := []byte(req.Image)
	req.Image = base64.StdEncoding.EncodeToString(textBytes)

	if err := u.todoRepository.CreateTodo(req); err != nil {
		return nil, err
	}

	filter := &models.TodosFilter{
		Order:   "created_at",
		OrderBy: "DESC",
	}

	res, err := u.todoRepository.ListTodos(filter)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *todoUsecase) ListTodos(filter *models.TodosFilter) ([]*models.TodoRes, error) {
	items, err := u.todoRepository.ListTodos(&models.TodosFilter{
		Order:   filter.Order,
		OrderBy: filter.OrderBy,
		Search:  filter.Search,
	})
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (u *todoUsecase) UpdateTodo(req *models.TodoReq, uid string) ([]*models.TodoRes, error) {
	textBytes := []byte(req.Image)
	req.Image = base64.StdEncoding.EncodeToString(textBytes)

	if err := u.todoRepository.UpdateTodos(req, uid); err != nil {
		return nil, err
	}

	filter := &models.TodosFilter{
		Order:   "created_at",
		OrderBy: "DESC",
	}

	res, err := u.todoRepository.ListTodos(filter)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *todoUsecase) TodoDone(uid string) ([]*models.TodoRes, error) {

	if err := u.todoRepository.TodoDone(uid); err != nil {
		return nil, err
	}

	filter := &models.TodosFilter{
		Order:   "created_at",
		OrderBy: "DESC",
	}

	res, err := u.todoRepository.ListTodos(filter)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *todoUsecase) RemoveTodo(uid string) ([]*models.TodoRes, error) {

	if err := u.todoRepository.RemoveTodo(uid); err != nil {
		return nil, err
	}

	filter := &models.TodosFilter{
		Order:   "created_at",
		OrderBy: "DESC",
	}

	res, err := u.todoRepository.ListTodos(filter)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func searchFilter(searchText string, todo *models.TodoRes) bool {
	return strings.Contains(strings.ToLower(todo.Title), strings.ToLower(searchText)) ||
		strings.Contains(strings.ToLower(todo.Description), strings.ToLower(searchText))
}

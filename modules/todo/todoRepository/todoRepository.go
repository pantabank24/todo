package todoRepository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pantabank24/todo/models"
)

type (
	TodoRepositoryService interface {
		CreateTodo(req *models.TodoReq) error
		ListTodos(filter *models.TodosFilter) ([]*models.TodoRes, error)
		UpdateTodos(req *models.TodoReq, uid string) error
		TodoDone(uid string) error
		RemoveTodo(uid string) error
	}

	todoRepository struct {
		Json string
	}
)

func NewTodoRepository(db *sqlx.DB) TodoRepositoryService {
	return &todoRepository{
		Json: "todos.json",
	}
}

func (r *todoRepository) CreateTodo(req *models.TodoReq) error {

	todos, err := ioutil.ReadFile(r.Json)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var items []models.TodoRes
	err = json.Unmarshal(todos, &items)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	items = append(items, models.TodoRes{
		ID:          uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Image:       req.Image,
		Status:      "IN_PROGRESS",
	})

	data, err := json.MarshalIndent(items, "", " ")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(r.Json, data, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil

}

func (r *todoRepository) ListTodos(filter *models.TodosFilter) ([]*models.TodoRes, error) {
	todos, err := ioutil.ReadFile(r.Json)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var items []*models.TodoRes
	err = json.Unmarshal(todos, &items)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var filteredItems []*models.TodoRes
	for _, item := range items {
		if searchFilter(filter.Search, item, "lists") {
			filteredItems = append(filteredItems, item)
		}
	}

	sort.Slice(filteredItems, func(i, j int) bool {
		if strings.ToLower(filter.Order) == "created_at" && strings.ToUpper(filter.OrderBy) == "ASC" {
			return filteredItems[i].CreatedAt.Before(filteredItems[j].CreatedAt)
		} else if strings.ToLower(filter.Order) == "title" && strings.ToUpper(filter.OrderBy) == "ASC" {
			return filteredItems[i].Title < filteredItems[j].Title
		} else if strings.ToLower(filter.Order) == "title" && strings.ToUpper(filter.OrderBy) == "DESC" {
			return filteredItems[i].Title > filteredItems[j].Title
		} else if strings.ToLower(filter.Order) == "status" && strings.ToUpper(filter.OrderBy) == "ASC" {
			return filteredItems[i].Status < filteredItems[j].Status
		} else if strings.ToLower(filter.Order) == "status" && strings.ToUpper(filter.OrderBy) == "DESC" {
			return filteredItems[i].Status > filteredItems[j].Status
		}
		return filteredItems[i].CreatedAt.After(filteredItems[j].CreatedAt)

	})

	return filteredItems, nil
}

func (r *todoRepository) UpdateTodos(req *models.TodoReq, uid string) error {
	todos, err := ioutil.ReadFile(r.Json)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var items []*models.TodoRes
	err = json.Unmarshal(todos, &items)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var update bool

	var filteredItems []*models.TodoRes
	for _, item := range items {
		if searchFilter(uid, item, "id") {
			if req.Description != "" {
				item.Description = req.Description
			}
			if req.Title != "" {
				item.Title = req.Title
			}
			if req.Image != "" {
				item.Image = req.Image
			}
			item.UpdatedAt = time.Now()
			filteredItems = append(filteredItems, item)
			update = true
		} else if item.ID != uid {
			filteredItems = append(filteredItems, item)
		}
	}

	data, err := json.MarshalIndent(filteredItems, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Json, data, 0644)
	if err != nil {
		return err
	}

	if !update {
		return fmt.Errorf("error: id not found : %s", uid)
	}

	return nil
}

func (r *todoRepository) TodoDone(uid string) error {
	todos, err := ioutil.ReadFile(r.Json)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var items []*models.TodoRes
	err = json.Unmarshal(todos, &items)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var done bool

	var filteredItems []*models.TodoRes
	for _, item := range items {
		if searchFilter(uid, item, "id") {
			item.Status = "COMPLETED"
			item.UpdatedAt = time.Now()
			filteredItems = append(filteredItems, item)
			done = true
		} else if item.ID != uid {
			filteredItems = append(filteredItems, item)
		}
	}

	data, err := json.MarshalIndent(filteredItems, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Json, data, 0644)
	if err != nil {
		return err
	}

	if !done {
		return fmt.Errorf("error: id not found : %s", uid)
	}

	return nil
}

func (r *todoRepository) RemoveTodo(uid string) error {
	todos, err := ioutil.ReadFile(r.Json)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var items []*models.TodoRes
	err = json.Unmarshal(todos, &items)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var remove bool

	var filteredItems []*models.TodoRes
	for _, item := range items {
		if item.ID == uid {
			remove = true
		} else {
			filteredItems = append(filteredItems, item)
		}
	}

	data, err := json.MarshalIndent(filteredItems, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Json, data, 0644)
	if err != nil {
		return err
	}

	if !remove {
		return fmt.Errorf("error: id not found : %s", uid)
	}

	return nil
}

func searchFilter(searchText string, todo *models.TodoRes, parameter string) bool {
	if parameter == "lists" {
		return strings.Contains(strings.ToLower(todo.Title), strings.ToLower(searchText)) ||
			strings.Contains(strings.ToLower(todo.Description), strings.ToLower(searchText))
	}

	if parameter == "id" {
		return strings.Contains(strings.ToLower(todo.ID), strings.ToLower(searchText))
	}

	return false

}

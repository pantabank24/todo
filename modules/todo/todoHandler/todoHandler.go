package todoHandler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pantabank24/todo/models"
	"github.com/pantabank24/todo/modules/todo/todoUsecase"
)

type (
	todoHandler struct {
		todoUsecase todoUsecase.TodoUsecaseService
	}
)

func NewTodoHandler(r fiber.Router, todoUsecase todoUsecase.TodoUsecaseService) {
	handler := &todoHandler{
		todoUsecase: todoUsecase,
	}
	r.Post("/", handler.CreateTodo)
	r.Get("/", handler.ListTodos)
	r.Patch("/:id", handler.UpdateTodo)
	r.Patch("/:id/completed", handler.TodoDone)
	r.Delete("/:id", handler.RemoveTodo)
}

func (h *todoHandler) CreateTodo(c *fiber.Ctx) error {
	req := new(models.TodoReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.ErrBadRequest.Code,
			"result":      nil,
		})
	}

	if req.Title == "" {
		return validateErrResponse("missing 'title'", c)
	}

	if len(req.Title) > 100 {
		return validateErrResponse("invalid 'title' (must not over 100 characters)", c)
	}

	res, err := h.todoUsecase.CreateTodo(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"result":      res,
	})
}

func (h *todoHandler) ListTodos(c *fiber.Ctx) error {
	filter := new(models.TodosFilter)
	if err := c.QueryParser(filter); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.ErrBadRequest.Code,
			"result":      nil,
		})
	}

	if filter.Order == "" || (strings.ToLower(filter.Order) != "title" && strings.ToLower(filter.Order) != "created_at" && strings.ToLower(filter.Order) != "status") {
		filter.Order = "created_at"
	}

	if filter.OrderBy == "" || (strings.ToUpper(filter.OrderBy) != "ASC" && strings.ToUpper(filter.OrderBy) != "DESC") {
		filter.OrderBy = "DESC"
	}

	res, err := h.todoUsecase.ListTodos(filter)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"result":      res,
	})
}

func (h *todoHandler) UpdateTodo(c *fiber.Ctx) error {
	uid := c.Params("id")

	req := new(models.TodoReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.ErrBadRequest.Code,
			"result":      nil,
		})
	}

	res, err := h.todoUsecase.UpdateTodo(req, uid)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"result":      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"result":      res,
	})
}

func (h *todoHandler) TodoDone(c *fiber.Ctx) error {
	uid := c.Params("id")

	res, err := h.todoUsecase.TodoDone(uid)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"result":      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"result":      res,
	})
}

func (h *todoHandler) RemoveTodo(c *fiber.Ctx) error {
	uid := c.Params("id")

	res, err := h.todoUsecase.RemoveTodo(uid)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"result":      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"result":      res,
	})
}

func validateErrResponse(msg string, c *fiber.Ctx) error {
	return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
		"status":      fiber.ErrBadRequest.Message,
		"status_code": fiber.ErrBadRequest.Code,
		"message":     msg,
	})
}

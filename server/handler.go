package server

import (
	"github.com/gofiber/fiber/v2"
	_todoHandler "github.com/pantabank24/todo/modules/todo/todoHandler"
	_todoRepository "github.com/pantabank24/todo/modules/todo/todoRepository"
	_todoUsecase "github.com/pantabank24/todo/modules/todo/todoUsecase"
)

func (s *Server) MapHandlers() error {
	v1 := s.App.Group("/v1")

	todoGroup := v1.Group("/todos")
	todoRepository := _todoRepository.NewTodoRepository(s.Db)
	todoUsecase := _todoUsecase.NewTodoUsecase(todoRepository)
	_todoHandler.NewTodoHandler(todoGroup, todoUsecase)

	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, endpoint not found",
			"result":      nil,
		})
	})

	return nil
}

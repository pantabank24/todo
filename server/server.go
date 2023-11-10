package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pantabank24/todo/config"
	"github.com/pantabank24/todo/pkg/utils"
)

type Server struct {
	App *fiber.App
	Cfg *config.Configs
	Db  *sqlx.DB
}

func NewServer(cfg *config.Configs) *Server {
	app := fiber.New()

	return &Server{
		App: app,
		Cfg: cfg,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatal(err.Error())
	}

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port

	log.Printf("server has been started on %s:%s", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatal(err.Error())
	}
}

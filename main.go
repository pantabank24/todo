package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pantabank24/todo/config"
	"github.com/pantabank24/todo/server"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err.Error())
	}
	cfg := new(config.Configs)

	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	s := server.NewServer(cfg)
	s.Start()
}

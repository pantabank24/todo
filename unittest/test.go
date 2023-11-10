package unittest

import (
	"os"

	"github.com/pantabank24/todo/config"
)

func NewTestConfig() *config.Configs {
	cfg := new(config.Configs)

	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	return cfg
}

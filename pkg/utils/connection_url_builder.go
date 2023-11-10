package utils

import (
	"errors"
	"fmt"

	"github.com/pantabank24/todo/config"
)

func ConnectionUrlBuilder(stuff string, cfg *config.Configs) (string, error) {
	var url string

	switch stuff {
	case "fiber":
		url = fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}
	return url, nil
}

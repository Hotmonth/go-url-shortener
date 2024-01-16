package main

import (
	"fmt"

	"github.com/Hotmonth/go-url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init loggger: slog

	// init: storage: sqlite

	// TODO: init router: chi, chi render

	// TODO: run server
}

package main

import (
	"log"

	"github.com/gobuffalo/packr"
	"github.com/t-0-m-1-3/todo/backend"
)

const (
	addressDefault = `127.0.0.1:8000`
)

func main() {

	config := backend.Config{
		Static:  packr.NewBox("frontend/dist/"),
		Address: addressDefault,
	}

	log.Print(backend.New(config).Start())
}

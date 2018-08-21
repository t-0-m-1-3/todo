package backend

import (
	"log"
	"net/http"
)

type Backend interface {
	Start() error
}

type backend struct {
	Config
}

func New(config Config) Backend {

	return &backend{Config: config}
}

func (b *backend) Start() error {

	http.Handle("/", http.FileServer(b.Static))
	log.Printf(`listening on %s`, b.Address)
	return http.ListenAndServe(b.Address, nil)
}

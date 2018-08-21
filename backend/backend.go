package backend

import (
	"log"
	"net/http"
)

func Start(config Config) error {

	http.Handle("/", http.FileServer(config.Static))
	log.Printf(`listening on %s`, config.Address)
	return http.ListenAndServe(config.Address, nil)
}

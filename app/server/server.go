package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/eqr/secrets/app/config"
)

type Server struct {
	router *chi.Mux
	url    string
}

func New(cfg config.Config) (*Server, error) {
	log.Println("starting server")

	db, err := bolt.Open(cfg.Database.Path, 0o600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("error creating db %v: %v", cfg.Database.Path, err.Error())
	}
	fmt.Println(db)

	router := chi.NewRouter()
	url := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Println("running server on ", url)

	return &Server{
		router: router,
		url:    url,
	}, nil
}

func (srv *Server) Start() error {
	if err := http.ListenAndServe(srv.url, srv.router); err != nil {
		log.Fatal("error running server: ", err.Error())
		return err
	}

	return nil
}

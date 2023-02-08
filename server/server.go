package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mxaxaxbx/go-rest-crud/database"
	"github.com/mxaxaxbx/go-rest-crud/repository"
	"github.com/mxaxaxbx/go-rest-crud/websocket"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required")
	}

	if config.DataBaseUrl == "" {
		return nil, errors.New("DataBaseUrl is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.config.DataBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	go b.hub.Run()
	repository.SetRepository(repo)

	log.Println("starting server on port", b.Config().Port)

	err = http.ListenAndServe(b.config.Port, b.router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Serv struct {
	log  *log.Logger
	serv *http.Server
}

func (s *Serv) ListenAndServe() error {
	return s.serv.ListenAndServe()
}

func CreateServer(logger *log.Logger) *Serv {
	// создаём новый роутер
	r := chi.NewRouter()

	// регистрируем в роутере эндпоинты
	r.Get("/", handlers.MainHandle)
	r.Post("/upload", handlers.UploadHandle)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Serv{
		log:  logger,
		serv: serv,
	}
}

package main

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {

	sm := mux.NewRouter()

	// register product handler
	NewProductHandler(sm)

	s := http.Server{
		Addr:         ":8001",
		Handler:      sm,
		TLSConfig:    &tls.Config{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	defer s.Shutdown(context.TODO())
	log.Ctx(context.TODO()).Info().Msg("Starting server...")
	err := s.ListenAndServe()
	if err != nil {
		log.Ctx(context.TODO()).Err(err).Msg("Server Failed")
	}
}

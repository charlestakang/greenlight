package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

)

const version = "1.0.0"

type config struct {
	port int
	env string
}
type application struct {
	config config
	logger *log.Logger
	version string
}
func main(){


	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger  := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
		version: version,
	}


	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}


}
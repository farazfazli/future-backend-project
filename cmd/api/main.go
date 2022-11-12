package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/farazfazli/future-backend-project/cmd/api/handlers"
	futuredb "github.com/farazfazli/future-backend-project/internal/db"
	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
}

func run() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	_, err = futuredb.NewQueries()
	if err != nil {
		return err
	}

	// Closes connection cleanly on CTRL+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Closing PostgreSQL DB...")
		if err := futuredb.CloseDB(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Println("Successfully closed DB!")
		os.Exit(0)
	}()

	mux := handlers.NewMux()
	api := &http.Server{
		Addr:           os.Getenv("HOST_URL"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10 * 1000, // 10KB
		Handler:        mux,
	}
	log.Printf("Listening on %s...\n", api.Addr)
	listenErr := api.ListenAndServe()
	if listenErr != nil {
		return listenErr
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
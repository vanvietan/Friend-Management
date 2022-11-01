package main

import (
	"fm/api/internal/config"
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	port := ":8000"
	r := chi.NewRouter()

	//Create database connection
	_, err := config.GetDatabaseConnection()
	if err != nil {
		log.Fatal("encountered error when create a db connection, error :%v", err)
	}

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}

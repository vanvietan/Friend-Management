package main

import (
	"fm/api/internal/api/router"
	"fm/api/internal/config"
	relationshipSvc "fm/api/internal/controller/relationship"
	userSvc "fm/api/internal/controller/user"
	"fm/api/internal/pkg"
	relationshipRepo "fm/api/internal/repository/relationship"
	userRepo "fm/api/internal/repository/user"
	"fmt"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	port := ":8000"
	r := chi.NewRouter()

	//Create database connection
	dbConn, err := config.GetDatabaseConnection()
	if err != nil {
		log.Fatal("encountered error when create a db connection, error :%v", err)
	}
	//ID sonyflake
	pkg.Init()
	userRepository := userRepo.New(dbConn)
	relationshipRepository := relationshipRepo.New(dbConn)
	userService := userSvc.New(userRepository)
	relationshipService := relationshipSvc.New(relationshipRepository)
	router.New(r, userService, relationshipService)

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}

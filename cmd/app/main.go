package main

import (
	"analytics_project/internal/db"
	"analytics_project/internal/handler"
	"analytics_project/internal/repository"
	"analytics_project/internal/service"
	"analytics_project/internal/util"
	"log"
	"net/http"
)

func main() {
	util.InitLogger()

	conn, err := db.ConnectMySQL()
	if err != nil {
		util.Error(err)
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	userRepo := repository.NewUserRepository(conn)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		util.Info("GET /users accessed")
		userHandler.GetAllUsersWithMenu(w, r)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		util.Info("GET /user accessed")
		userHandler.RunQuery(w, r)
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		util.Info("POST /token accessed")
		authHandler.CreateToken(w, r)
	})

	util.Info("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

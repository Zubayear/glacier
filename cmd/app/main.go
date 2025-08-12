package main

import (
	"database/sql"
	"glacier/internal/application/services"
	"glacier/internal/infrastructure/repository"
	"glacier/internal/infrastructure/server"
	http2 "glacier/internal/presentation/http"
	"log"
	"net/http"
)

func main() {
	// 1. Initialize Infrastructure components (the "dirty" details).
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Create the concrete repository, which implements the port.
	userRepo := repository.NewPgUserRepository(db)

	// 3. Create the use case, injecting the repository port.
	userService := services.NewUserService(userRepo)

	// 4. Create the handler, injecting the use case.
	userHandler := http2.NewUserHandler(userService)

	httpServer := server.NewServer(userHandler)
	httpServer.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", httpServer.GetRouter()))
}

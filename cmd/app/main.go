package main

import (
	"context"
	"glacier/internal/application/services"
	"glacier/internal/infrastructure/logger"
	"glacier/internal/infrastructure/repository"
	"glacier/internal/infrastructure/server"
	http2 "glacier/internal/presentation/http"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Config() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	const DATABASE_URL string = "postgres://postgres:password@localhost:5432/anubis_dev?"

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}

func main() {

	appLogger, err := logger.NewProductionLogger(true, "go-app")
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer appLogger.Logger.Sync()

	connPool, err := pgxpool.NewWithConfig(context.Background(), Config())

	// 2. Create the concrete repository, which implements the port.
	userRepo := repository.NewPgUserRepository(connPool)

	// 3. Create the use case, injecting the repository port.
	userService := services.NewUserService(userRepo, appLogger)

	// 4. Create the handler, injecting the use case.
	userHandler := http2.NewUserHandler(userService)

	httpServer := server.NewServer(userHandler)
	httpServer.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", httpServer.GetRouter()))
}

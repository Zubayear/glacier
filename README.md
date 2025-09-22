# Go Clean Architecture Template

[![Go Reference](https://pkg.go.dev/badge/github.com/Zubayear/glacier.svg)](https://pkg.go.dev/github.com/Zubayear/glacier)
[![Go Report Card](https://goreportcard.com/badge/github.com/Zubayear/glacier)](https://goreportcard.com/report/github.com/Zubayear/glacier)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

This repository provides a template for building Go applications using a Clean Architecture approach. It focuses on maintainable, testable, and scalable code by separating core business logic from external dependencies like databases, caches, loggers, and web frameworks.

## 🚀 Core Principles

This template adheres to the core principles of Clean Architecture

- The Dependency Rule: Dependencies must always point inwards. Inner circles should have no knowledge of outer circles. This is achieved in Go by defining interfaces (ports) in the inner layers and implementing them in the outer layers.
- Separation of Concerns: Each layer has a specific, well-defined responsibility.
- Testability: The core business logic can be tested in isolation, without needing to spin up a database or a web server.

## 📂 Project Structure

The project is organized into layers, each with a distinct role.

```
├── cmd
│   └── app
│       └── main.go           # Composition root: wires dependencies
├── internal                  # Core application code
│   ├── domain                # Business entities and rules
│   │   └── <component>.go
│   ├── application           # Use cases and interfaces (ports)
│   │   ├── ports             # Interfaces for external adapters
│   │   └── services          # Business logic (use cases)
│   ├── infrastructure       # External adapters (DB, cache, logger, server)
│   │   ├── repository        # DB implementations
│   │   └── <adapter>         # Other adapters (logger, cache, etc.)
│   └── presentation         # Controllers/handlers translating requests/responses
│       └── http              # HTTP-specific handlers
├── scaffold.sh               # Bash scaffolding script for components/adapters
├── go.mod
├── go.sum
└── README.md
```

- domain: Contains the pure, unadulterated business logic. This layer knows nothing about the outside world.
- application: Contains the use cases that orchestrate the domain entities. It defines the interfaces (ports) that the infrastructure layer must implement.
- infrastructure: Implements the interfaces defined in the application layer. This is where you'll find your database drivers, third-party libraries, and web framework code.
- presentation: The outermost layer. It adapts requests from external sources (like HTTP) into a format the application layer can understand, and then formats the application's response back to the client.
- cmd: The composition root. This is where all the dependencies are wired together. It's the only place that imports all other layers.

## 🛠️ Using This Template on GitHub

GitHub’s `Use this template` feature copies the code as–is.  
The module name in `go.mod` and the import paths will still be `github.com/Zubayear/glacier` until you change them.

### After creating your new repository:

1. Clone your new repository:

   ```bash
   git clone https://github.com/<your-username>/<your-new-repo>.git
   cd <your-new-repo>
   ```

2. Run the included init script to update the module name and imports:

   ```bash
   bash init.sh
   ```

   The script will prompt:

   ```
   Enter your module path (e.g., github.com/username/project):
   ```

   Type the correct module path for your new repo (for example, `github.com/janedoe/myapp`).  
   The script will:
   - Update the `module` directive in `go.mod`
   - Rewrite all internal import paths to use your module path
   - Run `go mod tidy` to clean up dependencies

3. Verify by opening `go.mod` — the module name should now match your repo URL.

## 🛠️ How to Run

1. Prerequisites:
   - Go (version 1.18 or higher)

2. Clone the repository:

```
git clone [your-repo-url]
cd [your-repo-name]
```

3. Get dependencies:

```
go mod tidy
```

4. Run the application:

```
go run cmd/app/main.go
```

The server will start on port `8080`.

5. Test the API:
   You can use a tool like curl to send a request to the `/users` endpoint.

```
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"John Doe", "email":"john.doe@example.com"}'
```

## 🛠️ Scaffold Script
The included scaffold.sh allows you to quickly generate and remove components and adapters.
```
# Create a new domain component
./scaffold.sh create component user

# Remove a domain component
./scaffold.sh remove component user

# Create a new infrastructure adapter
./scaffold.sh create adapter logger
./scaffold.sh create adapter cache
./scaffold.sh create adapter postgres

# Remove an adapter
./scaffold.sh remove adapter logger
```

## 📝 Adding a New Feature

This template is designed to make adding new features straightforward. Here is the typical workflow for adding a new API endpoint:

1. Start from the inside out:
   - Domain: Define any new entities or business rules.
   - Application: Create a new use case in the `application/services` package. This use case will contain the specific business logic for the new feature. Define any new ports (interfaces) in the `application/ports` package that the use case needs to interact with.
2. Work your way outwards:
   - Infrastructure: Create a new adapter in the `infrastructure` layer that implements the new port(s). For a database interaction, this would be a new repository implementation.
   - Presentation: Create a new handler in the `presentation/http` package that handles the incoming HTTP request and calls your new use case.
3. Wire it all up:
   - Composition Root (`main.go`): In `cmd/app/main.go`, create the concrete adapter(s), the new use case, and the new handler. Register the new endpoint with your server's router.

---
🔗 Example: Wiring Components
```go
package main

import (
	"log"
	"net/http"

	"glacier/internal/application/services"
	"glacier/internal/infrastructure/logger"
	"glacier/internal/infrastructure/repository"
	"glacier/internal/presentation/http"
)

func main() {
	// Infrastructure
	appLogger := logger.New()
	userRepo := repository.NewPGUserRepository()

	// Service
	userService := services.NewUserService(userRepo)

	// Handler
	userHandler := http.NewUserHandler(userService)

	// HTTP Route
	http.HandleFunc("/users/create", userHandler.CreateHandler)

	appLogger.Sugar().Info("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

```

# Go Clean Architecture Template
This repository provides a template for building Go applications using a Clean Architecture approach. The primary goal is to create a maintainable, testable, and scalable application by strictly separating the core business logic from external dependencies like databases, web frameworks, and file systems.

## 🚀 Core Principles
This template adheres to the core principles of Clean Architecture

* The Dependency Rule: Dependencies must always point inwards. Inner circles should have no knowledge of outer circles. This is achieved in Go by defining interfaces (ports) in the inner layers and implementing them in the outer layers.

* Separation of Concerns: Each layer has a specific, well-defined responsibility.

* Testability: The core business logic can be tested in isolation, without needing to spin up a database or a web server.

## 📂 Project Structure
The project is organized into layers, each with a distinct role.

```bash
├── cmd
│   └── app
│       └── main.go           # The application's entry point (the composition root).
├── internal                  # All core application code is here, not importable by other projects.
│   ├── application             # Business logic (use cases) and interfaces (ports).
│   │   ├── ports                 # Interfaces defining contracts for the outer layers.
│   │   └── services              # The "what" the application does (e.g., UserService).
│   ├── domain                  # Core business entities and rules.
│   │   └── user.go               # The "who" or "what" of your business (e.g., User struct).
│   ├── infrastructure          # External adapters for databases, servers, etc.
│   │   ├── repository            # Concrete database implementations.
│   │   └── server                # Web server setup and configuration.
│   ├── presentation            # Controllers/handlers that translate requests/responses.
│   │   └── http                  # HTTP-specific handlers (e.g., UserHandler).
├── go.mod                      # Go module file.
├── go.sum                      # Go module checksums.
└── README.md
```
* domain: Contains the pure, unadulterated business logic. This layer knows nothing about the outside world.

* application: Contains the use cases that orchestrate the domain entities. It defines the interfaces (ports) that the infrastructure layer must implement.

* infrastructure: Implements the interfaces defined in the application layer. This is where you'll find your database drivers, third-party libraries, and web framework code.

* presentation: The outermost layer. It adapts requests from external sources (like HTTP) into a format the application layer can understand, and then formats the application's response back to the client.

* cmd: The composition root. This is where all the dependencies are wired together. It's the only place that imports all other layers.

## 🛠️ How to Run
1. Prerequisites:
    - Go (version 1.18 or higher)
    - A PostgreSQL database running locally (or you can modify the code to use another database).

2. Clone the repository:
```bash
git clone [your-repo-url]
cd [your-repo-name]
```

3. Get dependencies:
```bash
go mod tidy
```

4. Run the application:

```bash
go run cmd/app/main.go
```

The server will start on port `8080`.

5. Test the API:
You can use a tool like curl to send a request to the `/users` endpoint.

```
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"John Doe", "email":"john.doe@example.com"}'
```

## 📝 Adding a New Feature
This template is designed to make adding new features straightforward. Here is the typical workflow for adding a new API endpoint:

1. Start from the inside out:

    * Domain: Define any new entities or business rules.

    * Application: Create a new use case in the application/services package. This use case will contain the specific business logic for the new feature. Define any new ports (interfaces) in the application/ports package that the use case needs to interact with.

2. Work your way outwards:

    * Infrastructure: Create a new adapter in the infrastructure layer that implements the new port(s). For a database interaction, this would be a new repository implementation.

    * Presentation: Create a new handler in the presentation/http package that handles the incoming HTTP request and calls your new use case.

3. Wire it all up:

    * Composition Root (main.go): In cmd/app/main.go, create the concrete adapter(s), the new use case, and the new handler. Register the new endpoint with your server's router.

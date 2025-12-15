# Copilot Instructions for goCode

## Project Overview
- This is a Go project (module: `github.com/NaveenChander/GoFace`) with its main logic in the `simulator/` directory.
- The entry point is `simulator/main.go`.
- The project demonstrates HTTP requests and PostgreSQL database interactions.

## Architecture & Patterns
- All logic currently resides in a single file: `simulator/main.go`.
- Key functions:
  - `main()`: Entry point, prints a message and demonstrates an HTTP GET request.
  - `makeGetRequest()`: Makes a sample HTTP GET request and prints the response.
  - `createPostgresConnection()`: Connects to a PostgreSQL database using a connection string.
  - `insertRow()`: Inserts a row into a specified table using dynamic columns and values.
  - `joinColumns()` and `stringJoin()`: Utility functions for SQL query construction.
- Uses the `github.com/lib/pq` driver for PostgreSQL.

## Developer Workflows
- **Build:** Run `go build ./simulator` from the project root to build the simulator.
- **Run:** Execute `go run simulator/main.go` to run the main program.
- **Dependencies:** Managed via Go modules (`go.mod`). Add new dependencies with `go get <package>`.
- **Database:** Requires a valid PostgreSQL connection string for DB operations. No migrations or schema files are present.

## Conventions & Practices
- All code is in `simulator/main.go`. No package structure beyond this yet.
- HTTP and DB logic are separated into functions for clarity.
- SQL queries are constructed dynamically; ensure column/value alignment.
- Error handling is basic: errors are printed to stdout.

## Integration Points
- **External HTTP API:** Example GET request to `jsonplaceholder.typicode.com`.
- **PostgreSQL:** Uses `github.com/lib/pq` for DB access. Connection string must be provided at runtime.

## Extending the Project
- Add new features as separate functions in `simulator/main.go` or refactor into packages as the codebase grows.
- Follow the pattern of separating concerns (HTTP, DB, utilities).

## Example Commands
- Build: `go build ./simulator`
- Run: `go run simulator/main.go`
- Add dependency: `go get github.com/some/package`

---

_If you add new directories, files, or workflows, update this file to keep AI agents productive._

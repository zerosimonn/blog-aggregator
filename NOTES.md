# Gator Notes

## Stack

PostgreSQL  = database server (also called "Postgres")
Goose       = database migration tool (applies versioned schema changes up/down)
sqlc        = code generator (reads your SQL, outputs type-safe Go functions)
SQL         = query language for relational databases (declarative: you say what, not how)
Go          = compiled, statically-typed programming language (application code)
VS Code     = code editor (note, could be any editor; not part of the stack)
psql        = command-line client for Postgres (for ad-hoc queries and inspection)
database/sql = Go's standard library interface for talking to any SQL database (the generic socket)
lib/pq      = Postgres driver that plugs into database/sql (the specific plug)
google/uuid = Go library for generating UUIDs in your app code
~/.gatorconfig.json = your CLI's config file (stores DB URL and current user)

You write:                 Tool processes:           Result:
migrations (.sql)    ──►   Goose         ──►   schema in Postgres
queries (.sql)       ──►   sqlc          ──►   Go functions
Go code (main.go)    ──►   go build      ──►   gator binary
                                               │
                                               ▼
                            gator ──► database/sql ──► lib/pq ──► Postgres

## Project layout

gator/                          ← project root (run sqlc & goose from here)
├── main.go                     ← entry point, opens DB, builds state, dispatches commands
├── commands.go                 ← command registry (maps "register" → handlerRegister, etc.)
├── handler_user.go             ← one file per command group (register, login, ...)
├── sqlc.yaml                   ← sqlc config (where to read SQL, where to write Go)
├── go.mod / go.sum             ← Go module + dependency lockfile
│
├── internal/                   ← private packages (not importable by other modules)
│   ├── config/                 ← reads/writes ~/.gatorconfig.json
│   │   └── config.go
│   └── database/               ← sqlc OUTPUT — never edit by hand
│       ├── db.go               ← generated: connection wrapper
│       ├── models.go           ← generated: Go structs for each table
│       └── users.sql.go        ← generated: Go funcs for users.sql queries
│
└── sql/                        ← all SQL lives here
    ├── schema/                 ← Goose migrations (numbered, with up/down)
    │   └── 001_users.sql
    └── queries/                ← sqlc INPUT (your hand-written queries)
        └── users.sql

## SQLC

- Config lives in `sqlc.yaml` at project root. Run `sqlc generate` from there.
- Schema dir = same files Goose uses; SQLC ignores `down` migrations automatically.
- Query annotations on the line above the SQL:
  - `-- name: CreateUser :one`   → returns one row
  - `-- name: GetUsers :many`    → returns a slice
  - `-- name: DeleteUser :exec`  → returns nothing (just runs)
- `$1, $2, ...` are positional params → become Go function args in order.
- `RETURNING *` is what makes `:one` actually give you the row back.
- Generated code lands in `internal/database/`. Never edit it; regenerate.

## Goose

- Migrations live in `sql/schema/`, numbered: `001_users.sql`, `002_...`.
- Each file has `-- +goose Up` and `-- +goose Down` sections.
- Before testing/submitting: `goose down` then `goose up` for a clean DB.
- cat ~/.gatorconfig.json // to retrieve the Postgres connection string. 
- postgres://postgres:postgres@localhost:5432/gator?sslmode=disable
- goose -dir sql/schema postgres <connection string> up
- goose -dir sql/schema postgres postgres://postgres:postgres@localhost:5432/gator?sslmode=disable up

## Postgres driver (lib/pq)

- Install: `go get github.com/lib/pq`
- Import for side effects only:
  ```go
  import _ "github.com/lib/pq"
  ```
- The `_` means: register the driver, don't expose its symbols.

## database/sql

- Open a connection:
  ```go
  db, err := sql.Open("postgres", dbURL)
  ```
- `"postgres"` is the driver name registered by `lib/pq`.
- Wrap with SQLC's generated constructor:
  ```go
  dbQueries := database.New(db)
  ```

## Git work reminder
git checkout -b feature-branchname
git add .
git commit -m "wip: halfway there"
git commit -m "done"

git checkout main
git merge --squash feature-brachnname

git commit -m "added feature .."
git push origin main //only push a commit 

git status // staging area
git log --oneline --graph --all

## UUIDs (google/uuid)

- Install: `go get github.com/google/uuid`
- Generate a new ID: `uuid.New()` → returns `uuid.UUID`.
- App-side IDs (vs Postgres-side) → portable, no DB round-trip to get the ID.

## Context

- Every SQLC-generated query takes a `context.Context` first arg.
- For CLI commands with no deadline: `context.Background()`.

## Project layout (Gator)

```
sqlc.yaml
sql/
  schema/    ← Goose migrations
  queries/   ← SQLC query files (*.sql)
internal/
  database/  ← SQLC-generated Go (do not edit)
  config/    ← ~/.gatorconfig.json read/write
main.go
commands.go
handler_*.go
```

## Common workflow

1. Add migration file in `sql/schema/`
2. `goose up`
3. Write query in `sql/queries/*.sql`
4. `sqlc generate`
5. Use `state.db.YourQuery(ctx, args)` in a handler
6. Test: `go run . <command> <args>`

## Example code addition

Full workflow:

sql/queries/users.sql

add the delete-all-users query
sqlc generate

generates the Go DB method in internal/database/
handler_reset.go

write handlerReset
call s.db.ResetUsers(context.Background())
return error on failure, print success on success
commands.go

add "reset" to the command registry
point it at handlerReset
main.go

parses CLI args
looks up the command in the registry
runs the matching handler
So main.go is the entry point, commands.go is the routing table, and the handler does the work.

Tiny mental model:

go run . reset
   -> main.go
   -> commands.go finds "reset"
   -> handler_reset.go runs
   -> calls generated DB code
   -> SQL deletes all rows from users
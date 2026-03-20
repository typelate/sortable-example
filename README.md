# Sortable Example

A drag-and-drop task list built with Go, htmx, and SortableJS. Tasks within a list can be reordered by dragging, and the new order is persisted to PostgreSQL.

## Stack

- [muxt](https://github.com/typelate/muxt) — template-driven HTTP routing
- [htmx](https://htmx.org) + [SortableJS](https://sortablejs.github.io/Sortable/) — drag-and-drop reordering without a JS framework
- [sqlc](https://sqlc.dev) — type-safe SQL queries
- [pgx](https://github.com/jackc/pgx) — PostgreSQL driver

## Running

Set the `DATABASE_URL` environment variable and run:

```sh
go run github.com/typelate/sortable-example/cmd/server
```

The server listens on `PORT` (default `8080`).

[Goose](https://github.com/pressly/goose)-like migrations run automatically on startup.

## Testing

Unit tests (no database required):

```sh
go test ./...
```

Integration tests in `internal/database` require a running PostgreSQL instance reachable via `DATABASE_URL`. They are skipped automatically when the database is unavailable.

## License

MIT
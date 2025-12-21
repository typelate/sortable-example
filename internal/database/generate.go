package database

//go:generate sqlc generate
//go:generate querier-interface queries.sql
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fake/tx.go --fake-name Tx github.com/jackc/pgx/v5.Tx
//counterfeiter:generate -o internal/fake/transaction_manager.go --fake-name TransactionManager . Caller

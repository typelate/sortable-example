package database

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/typelate/loosey"
)

func Setup(ctx context.Context) (*pgxpool.Pool, error) {
	db, err := Connect(ctx, pgxpool.New)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := RunMigrations(ctx, db); err != nil {
		return nil, err
	}
	return db, nil
}

func Name() (string, error) {
	value, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		return "", fmt.Errorf("DATABASE_NAME environment variable not set")
	}
	return value, nil
}

func URL() (string, error) {
	value, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return "", fmt.Errorf("DATABASE_URL environment variable not set")
	}
	return value, nil
}

//go:embed schema/*.sql
var migrations embed.FS

func RunMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	schemaDir, err := fs.Sub(migrations, "schema")
	if err != nil {
		return err
	}
	db := stdlib.OpenDBFromPool(pool)
	defer func() {
		_ = db.Close()
	}()
	m, err := loosey.NewPostgres(ctx, db, schemaDir)
	if err != nil {
		return err
	}
	if _, err := m.Up(ctx); err != nil {
		return err
	}
	return nil
}

const ConnectPingTimeout = 5 * time.Second

type Pinger interface {
	DBTX
	Ping(context.Context) error
}

type NewFunc[DB Pinger] func(ctx context.Context, dbURL string) (DB, error)

func Connect[DB Pinger](ctx context.Context, newDB NewFunc[DB]) (DB, error) {
	databaseURL, err := URL()
	if err != nil {
		var zero DB
		return zero, fmt.Errorf("could not determine database URL: %w", err)
	}
	dbConn, err := newDB(ctx, databaseURL)
	if err != nil {
		var zero DB
		return zero, fmt.Errorf("failed to create database pool: %w", err)
	}
	pingCtx, cancel := context.WithTimeout(ctx, ConnectPingTimeout)
	defer cancel()
	if err := dbConn.Ping(pingCtx); err != nil {
		return dbConn, fmt.Errorf("failed to ping database: %w", err)
	}
	return dbConn, nil
}

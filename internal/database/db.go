package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DSN struct {
	user     string
	password string
	host     string
	port     int
	dbName   string
	sslMode  string
	maxConns int
}

func NewPool() *pgxpool.Pool {

	db, err := pgxpool.NewWithConfig(context.Background(), config())
	if err != nil {
		log.Fatal(fmt.Errorf("error creating pool with config, %w", err))
	}

	return db

}

func connString() string {
	dsn := DSN{
		user:     envOrDefault("POSTGRES_USER", "admin"),
		password: envOrDefault("POSTGRES_PASSWORD", "123"),
		host:     envOrDefault("POSTGRES_HOSTNAME", "localhost"),
		port:     5432,
		dbName:   envOrDefault("POSTGRES_DB", "rinha"),
		sslMode:  "disable",
		maxConns: 12,
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&pool_max_conns=%d",
		dsn.user,
		dsn.password,
		dsn.host,
		dsn.port,
		dsn.dbName,
		dsn.sslMode,
		dsn.maxConns,
	)

	return connString
}

func config() *pgxpool.Config {

	connString := connString()
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal("error parsing db config ", err)
	}

	config.MinConns = 40
	config.MaxConns = 50
	config.MaxConnIdleTime = time.Minute * 3

	return config
}

func envOrDefault(name, fallback string) string {

	env := os.Getenv(name)
	if env == "" {
		env = fallback
	}

	return env
}

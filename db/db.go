package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"os"
	"strings"
)

var pool *pgxpool.Pool

// Init starts a database connection,
// populates the database with the init sql file,
// and initializes connection pool variable
func Init() error {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	pool = conn

	if err = loadSQLFile(pool, "db/init.sql"); err != nil {
		return err
	}

	return nil
}

// Pool returns the connection pool pointer
func Pool() *pgxpool.Pool {
	return pool
}

// Helper function to load sql file
func loadSQLFile(pool *pgxpool.Pool, filepath string) error {
	// Read sql file into bytes
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Begin transaction
	tx, err := pool.Begin(context.Background())
	if err != nil {
		return err
	}
	// Defer transaction rollback, only called in case of failure
	defer tx.Rollback(context.Background())

	// Execute each query in sql file
	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(context.Background(), q); err != nil {
			return err
		}
	}

	return tx.Commit(context.Background())
}

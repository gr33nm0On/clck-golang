package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Failed DB connection: %v", err)
	}

	return pool
}

func Save(ctx context.Context, pool *pgxpool.Pool, origin_url string, short_url string) {
	var url string

	err := pool.QueryRow(ctx, "SELECT origin_url FROM urls WHERE origin_url = $1", origin_url).Scan(&url)

	fmt.Print("urlq: ", url)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Printf("Query error: %v", err)
			return
		}
		log.Printf("URL not exists in DB")
	}

	_, err = pool.Exec(ctx, `
			INSERT INTO urls (origin_url, short_url) 
			VALUES ($1, $2)
			ON CONFLICT (origin_url) 
			DO UPDATE SET origin_url = EXCLUDED.origin_url
			RETURNING short_url;
		`, origin_url, short_url)

	if err != nil {
		log.Fatal(err)
	}
}

func InitDB(ctx context.Context, pool *pgxpool.Pool) error {
	initCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
	CREATE TABLE IF NOT EXISTS urls (
	    id serial PRIMARY KEY,
	    origin_url text NOT NULL UNIQUE,
	    short_url text NOT NULL UNIQUE
	)
	`

	_, err := pool.Exec(initCtx, query)
	return err
}

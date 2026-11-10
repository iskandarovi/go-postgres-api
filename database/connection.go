package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB(ctx context.Context) (*pgxpool.Pool, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	

	dbURL := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	pingErr := pool.Ping(ctx)
	if pingErr != nil {
		pool.Close() 
		return nil, pingErr
	}
	return pool, nil
}
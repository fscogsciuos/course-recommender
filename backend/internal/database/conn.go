package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	dbb "github.com/tomhaerter/course-recommender/internal/db"
	"os"
)

type Db struct {
	Db *pgxpool.Pool
	*dbb.Queries
}

func NewClient() *Db {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}

	return &Db{
		Db:      pool,
		Queries: dbb.New(pool),
	}
}

package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

func NewPgxConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
}

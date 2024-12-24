package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"url-shortener-api/internal/entity"
)

type URLRepository struct {
	conn *pgx.Conn
}

func (r *URLRepository) Save(ctx context.Context, url *entity.URL) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO urls (id, target_link, hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", url.ID, url.TargetLink, url.Hash, url.CreatedAt, url.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *URLRepository) FindByHash(ctx context.Context, hash string) (entity.URL, error) {
	var url entity.URL
	err := r.conn.QueryRow(ctx, "SELECT id, target_link, hash, created_at, updated_at FROM urls WHERE hash = $1", hash).Scan(&url.ID, &url.TargetLink, &url.Hash, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		return entity.URL{}, err
	}
	return url, nil
}

func NewURLRepository(conn *pgx.Conn) *URLRepository {
	return &URLRepository{conn: conn}
}

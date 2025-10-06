package database

import (
	"context"
	"database/sql"

	"github.com/farinas09/feeds-api/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertFeed(ctx context.Context, feed *models.Feed) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO feeds (id, title, description) VALUES ($1, $2, $3)", feed.Id, feed.Title, feed.Description)
	return err
}

func (r *PostgresRepository) ListFeeds(ctx context.Context) ([]*models.Feed, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, description, created_at FROM feeds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feeds := []*models.Feed{}

	for rows.Next() {
		var feed models.Feed
		err := rows.Scan(&feed.Id, &feed.Title, &feed.Description, &feed.CreatedAt)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, &feed)
	}
	return feeds, nil
}

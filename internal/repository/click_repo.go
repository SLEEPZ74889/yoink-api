package repository

import (
	"context"
	"fmt"

	"github.com/SLEEPZ74889/yoink-api/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClickRepository struct {
	db *pgxpool.Pool
}

func NewClickRepository(db *pgxpool.Pool) *ClickRepository {
	return &ClickRepository{db: db}
}

func (r *ClickRepository) Create(ctx context.Context, click *model.Click) error {
	query := `
		INSERT INTO clicks (link_id, referrer, browser, os, device, ip_hash)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, clicked_at
	`
	err := r.db.QueryRow(ctx, query,
		click.LinkID,
		click.Referrer,
		click.Browser,
		click.OS,
		click.Device,
		click.IPHash,
	).Scan(&click.ID, &click.ClickedAt)

	if err != nil {
		return fmt.Errorf("repository.ClickCreate: %w", err)
	}
	return nil
}

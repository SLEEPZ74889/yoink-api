package repository

import (
	"context"
	"fmt"

	"github.com/SLEEPZ74889/yoink-api/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepository struct {
	db *pgxpool.Pool
}

func NewLinkRepository(db *pgxpool.Pool) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(ctx context.Context, link *model.Link) error {
	query := `
		INSERT INTO links (slug, original_url, title, active)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	err := r.db.QueryRow(ctx, query,
		link.Slug,
		link.OriginalURL,
		link.Title,
		link.Active,
	).Scan(&link.ID, &link.CreatedAt)

	if err != nil {
		return fmt.Errorf("repository.Create: %w", err)
	}
	return nil
}

func (r *LinkRepository) FindBySlug(ctx context.Context, slug string) (*model.Link, error) {
	query := `
		SELECT id, slug, original_url, title, active, created_at
		FROM links
		WHERE slug = $1
	`
	link := &model.Link{}
	err := r.db.QueryRow(ctx, query, slug).Scan(
		&link.ID,
		&link.Slug,
		&link.OriginalURL,
		&link.Title,
		&link.Active,
		&link.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("repository.FindBySlug: %w", err)
	}
	return link, nil
}

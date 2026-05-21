package service

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"strings"

	"github.com/SLEEPZ74889/yoink-api/internal/model"
	"github.com/SLEEPZ74889/yoink-api/internal/repository"
)

type LinkService struct {
	repo *repository.LinkRepository
}

func NewLinkService(repo *repository.LinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) CreateLink(ctx context.Context, req *model.CreateLinkRequest) (*model.Link, error) {
	if strings.TrimSpace(req.URL) == "" {
		return nil, fmt.Errorf("url darf nicht leer sein")
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		return nil, fmt.Errorf("ungültige url: %w", err)
	}

	slug := strings.TrimSpace(req.Slug)
	if slug == "" {
		slug = generateSlug(6)
	}

	link := &model.Link{
		Slug:        slug,
		OriginalURL: req.URL,
		Title:       req.Title,
		Active:      true,
	}

	if err := s.repo.Create(ctx, link); err != nil {
		return nil, fmt.Errorf("service.CreateLink: %w", err)
	}

	return link, nil
}

func (s *LinkService) GetLinkBySlug(ctx context.Context, slug string) (*model.Link, error) {
	link, err := s.repo.FindBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("service.GetLinkBySlug: %w", err)
	}

	if !link.Active {
		return nil, fmt.Errorf("link ist deaktiviert")
	}

	return link, nil
}

func generateSlug(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

package service

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"strings"

	"github.com/SLEEPZ74889/yoink-api/internal/model"
)

// Interfaces — der Service hängt von diesen ab, nicht von konkreten Typen.
// Das ermöglicht Mocks in Tests.
type LinkRepo interface {
	Create(ctx context.Context, link *model.Link) error
	FindBySlug(ctx context.Context, slug string) (*model.Link, error)
}

type ClickRepo interface {
	Create(ctx context.Context, click *model.Click) error
}

type LinkService struct {
	links  LinkRepo
	clicks ClickRepo
}

func NewLinkService(links LinkRepo, clicks ClickRepo) *LinkService {
	return &LinkService{links: links, clicks: clicks}
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

	if err := s.links.Create(ctx, link); err != nil {
		return nil, fmt.Errorf("service.CreateLink: %w", err)
	}

	return link, nil
}

func (s *LinkService) GetLinkBySlug(ctx context.Context, slug string) (*model.Link, error) {
	link, err := s.links.FindBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("link nicht gefunden")
	}

	if !link.Active {
		return nil, fmt.Errorf("link ist deaktiviert")
	}

	return link, nil
}

// TrackClick speichert einen Klick asynchron — wird vom Handler per goroutine aufgerufen.
func (s *LinkService) TrackClick(ctx context.Context, linkID, referrer, browser, os, device, ipHash string) {
	click := &model.Click{
		LinkID:   linkID,
		Referrer: referrer,
		Browser:  browser,
		OS:       os,
		Device:   device,
		IPHash:   ipHash,
	}
	s.clicks.Create(ctx, click)
}

func generateSlug(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

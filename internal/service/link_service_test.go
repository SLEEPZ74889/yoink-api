package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/SLEEPZ74889/yoink-api/internal/model"
)

// --- Mock Implementations ---

type mockLinkRepo struct {
	store map[string]*model.Link
	err   error
}

func newMockLinkRepo() *mockLinkRepo {
	return &mockLinkRepo{store: make(map[string]*model.Link)}
}

func (m *mockLinkRepo) Create(ctx context.Context, link *model.Link) error {
	if m.err != nil {
		return m.err
	}
	link.ID = "test-uuid-1234"
	m.store[link.Slug] = link
	return nil
}

func (m *mockLinkRepo) FindBySlug(ctx context.Context, slug string) (*model.Link, error) {
	if m.err != nil {
		return nil, m.err
	}
	link, ok := m.store[slug]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return link, nil
}

type mockClickRepo struct {
	called int
}

func (m *mockClickRepo) Create(ctx context.Context, click *model.Click) error {
	m.called++
	return nil
}

// --- Tests ---

func TestCreateLink_GeneratesSlug(t *testing.T) {
	svc := NewLinkService(newMockLinkRepo(), &mockClickRepo{})

	link, err := svc.CreateLink(context.Background(), &model.CreateLinkRequest{
		URL: "https://google.com",
	})

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if link.Slug == "" {
		t.Error("expected slug to be generated, got empty string")
	}
	if len(link.Slug) != 6 {
		t.Errorf("expected slug length 6, got %d", len(link.Slug))
	}
}

func TestCreateLink_CustomSlug(t *testing.T) {
	svc := NewLinkService(newMockLinkRepo(), &mockClickRepo{})

	link, err := svc.CreateLink(context.Background(), &model.CreateLinkRequest{
		URL:  "https://google.com",
		Slug: "meinslug",
	})

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if link.Slug != "meinslug" {
		t.Errorf("expected slug 'meinslug', got '%s'", link.Slug)
	}
}

func TestCreateLink_EmptyURL(t *testing.T) {
	svc := NewLinkService(newMockLinkRepo(), &mockClickRepo{})

	_, err := svc.CreateLink(context.Background(), &model.CreateLinkRequest{URL: ""})

	if err == nil {
		t.Error("expected error for empty URL, got nil")
	}
}

func TestCreateLink_InvalidURL(t *testing.T) {
	svc := NewLinkService(newMockLinkRepo(), &mockClickRepo{})

	_, err := svc.CreateLink(context.Background(), &model.CreateLinkRequest{
		URL: "das-ist-keine-url",
	})

	if err == nil {
		t.Error("expected error for invalid URL, got nil")
	}
}

func TestCreateLink_SetsActiveTrue(t *testing.T) {
	svc := NewLinkService(newMockLinkRepo(), &mockClickRepo{})

	link, err := svc.CreateLink(context.Background(), &model.CreateLinkRequest{
		URL: "https://example.com",
	})

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if !link.Active {
		t.Error("expected link.Active to be true")
	}
}

func TestGetLinkBySlug_Success(t *testing.T) {
	repo := newMockLinkRepo()
	repo.store["abc123"] = &model.Link{
		ID:          "uuid",
		Slug:        "abc123",
		OriginalURL: "https://example.com",
		Active:      true,
	}
	svc := NewLinkService(repo, &mockClickRepo{})

	link, err := svc.GetLinkBySlug(context.Background(), "abc123")

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if link.OriginalURL != "https://example.com" {
		t.Errorf("expected URL 'https://example.com', got '%s'", link.OriginalURL)
	}
}

func TestGetLinkBySlug_NotFound(t *testing.T) {
	svc := NewLinkService(newMockLinkRepo(), &mockClickRepo{})

	_, err := svc.GetLinkBySlug(context.Background(), "gibts-nicht")

	if err == nil {
		t.Error("expected error for unknown slug, got nil")
	}
}

func TestGetLinkBySlug_DeactivatedLink(t *testing.T) {
	repo := newMockLinkRepo()
	repo.store["deaktiv"] = &model.Link{
		ID:          "uuid",
		Slug:        "deaktiv",
		OriginalURL: "https://example.com",
		Active:      false,
	}
	svc := NewLinkService(repo, &mockClickRepo{})

	_, err := svc.GetLinkBySlug(context.Background(), "deaktiv")

	if err == nil {
		t.Error("expected error for deactivated link, got nil")
	}
}

func TestCreateLink_DBError(t *testing.T) {
	repo := newMockLinkRepo()
	repo.err = fmt.Errorf("db connection lost")
	svc := NewLinkService(repo, &mockClickRepo{})

	_, err := svc.CreateLink(context.Background(), &model.CreateLinkRequest{
		URL: "https://example.com",
	})

	if err == nil {
		t.Error("expected error when DB fails, got nil")
	}
}

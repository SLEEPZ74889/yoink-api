package model

import "time"

type Link struct {
	ID          string
	Slug        string
	OriginalURL string
	Title       string
	Active      bool
	CreatedAt   time.Time
	ExpiresAt   *time.Time
}

type CreateLinkRequest struct {
	URL   string `json:"url"`
	Slug  string `json:"slug,omitempty"`
	Title string `json:"title,omitempty"`
}

type CreateLinkResponse struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	CreatedAt   string `json:"created_at"`
}

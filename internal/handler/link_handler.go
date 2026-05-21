package handler

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/SLEEPZ74889/yoink-api/internal/model"
	"github.com/SLEEPZ74889/yoink-api/internal/service"
	"github.com/go-chi/chi/v5"
)

type LinkHandler struct {
	service *service.LinkService
}

func NewLinkHandler(service *service.LinkService) *LinkHandler {
	return &LinkHandler{service: service}
}

func (h *LinkHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "ungültiger request body", http.StatusBadRequest)
		return
	}

	link, err := h.service.CreateLink(r.Context(), &req)
	if err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := model.CreateLinkResponse{
		ID:          link.ID,
		Slug:        link.Slug,
		ShortURL:    "https://yoink.link/" + link.Slug,
		OriginalURL: link.OriginalURL,
		CreatedAt:   link.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *LinkHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	link, err := h.service.GetLinkBySlug(r.Context(), slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Click-Tracking asynchron — Redirect soll sofort passieren
	go h.service.TrackClick(
		context.Background(),
		link.ID,
		r.Header.Get("Referer"),
		parseBrowser(r.Header.Get("User-Agent")),
		parseOS(r.Header.Get("User-Agent")),
		parseDevice(r.Header.Get("User-Agent")),
		hashIP(getIP(r)),
	)

	http.Redirect(w, r, link.OriginalURL, http.StatusFound)
}

func jsonError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func getIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return strings.Split(forwarded, ",")[0]
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func hashIP(ip string) string {
	h := sha256.Sum256([]byte(ip))
	return fmt.Sprintf("%x", h)
}

func parseBrowser(ua string) string {
	switch {
	case strings.Contains(ua, "Edg"):
		return "Edge"
	case strings.Contains(ua, "Chrome"):
		return "Chrome"
	case strings.Contains(ua, "Firefox"):
		return "Firefox"
	case strings.Contains(ua, "Safari"):
		return "Safari"
	default:
		return "Other"
	}
}

func parseOS(ua string) string {
	switch {
	case strings.Contains(ua, "Windows"):
		return "Windows"
	case strings.Contains(ua, "Mac OS"):
		return "macOS"
	case strings.Contains(ua, "Linux"):
		return "Linux"
	case strings.Contains(ua, "Android"):
		return "Android"
	case strings.Contains(ua, "iPhone"), strings.Contains(ua, "iPad"):
		return "iOS"
	default:
		return "Other"
	}
}

func parseDevice(ua string) string {
	switch {
	case strings.Contains(ua, "Mobile"), strings.Contains(ua, "Android"), strings.Contains(ua, "iPhone"):
		return "mobile"
	case strings.Contains(ua, "Tablet"), strings.Contains(ua, "iPad"):
		return "tablet"
	default:
		return "desktop"
	}
}

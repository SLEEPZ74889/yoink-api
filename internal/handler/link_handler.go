package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SLEEPZ74889/yoink-api/internal/model"
	"github.com/SLEEPZ74889/yoink-api/internal/service"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ungültiger request body"})
		return
	}

	link, err := h.service.CreateLink(r.Context(), &req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
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

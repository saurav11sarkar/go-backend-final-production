package link

import (
	"encoding/json"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/http"
)

type Handler struct{ s *Service }

func NewHandler(s *Service) *Handler { return &Handler{s} }
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	l, err := h.s.Create(r.Context(), in.Title, in.URL)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	utils.JSON(w, 201, "link created", l)
}
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.s.List(r.Context())
	if err != nil {
		utils.Error(w, 500, "failed to fetch links")
		return
	}
	utils.JSON(w, 200, "links fetched", list)
}

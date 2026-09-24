package user

import (
	"encoding/json"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/middleware"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/http"
)

type Handler struct{ service *Service }

func NewHandler(s *Service) *Handler { return &Handler{service: s} }
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	id, _ := middleware.CurrentUser(r)
	u, err := h.service.Me(r.Context(), id)
	if err != nil {
		utils.Error(w, 404, "user not found")
		return
	}
	utils.JSON(w, 200, "profile fetched", u)
}
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	q := utils.ParseQuery(r.URL.Query())
	list, total, err := h.service.List(r.Context(), q)
	if err != nil {
		utils.Error(w, 500, "failed to fetch users")
		return
	}
	utils.JSON(w, 200, "users fetched", map[string]any{"items": list, "page": q.Page, "limit": q.Limit, "total": total})
}
func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var in struct {
		FullName string `json:"fullName" validate:"required,min=2,max=100"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	id, _ := middleware.CurrentUser(r)
	u, err := h.service.UpdateProfile(r.Context(), id, in.FullName)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	utils.JSON(w, 200, "profile updated", u)
}

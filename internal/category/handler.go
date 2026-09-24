package category

import (
	"encoding/json"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/http"
)

type Handler struct{ s *Service }

func NewHandler(s *Service) *Handler { return &Handler{s} }
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Name string `json:"name" validate:"required,min=2,max=100"`
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	c, err := h.s.Create(r.Context(), in.Name)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	utils.JSON(w, 201, "category created", c)
}
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.s.List(r.Context(), utils.ParseQuery(r.URL.Query()))
	if err != nil {
		utils.Error(w, 500, "failed to fetch categories")
		return
	}
	utils.JSON(w, 200, "categories fetched", list)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Name string `json:"name" validate:"required,min=2,max=100"`
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	c, err := h.s.Update(r.Context(), r.PathValue("id"), in.Name)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	utils.JSON(w, 200, "category updated", c)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.s.Delete(r.Context(), r.PathValue("id")); err != nil {
		utils.Error(w, 409, "category cannot be deleted while products are using it")
		return
	}
	utils.JSON(w, 200, "category deleted", nil)
}

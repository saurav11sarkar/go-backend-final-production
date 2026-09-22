package product

import (
	"encoding/json"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/http"
)

type Handler struct{ s *Service }

func NewHandler(s *Service) *Handler { return &Handler{s} }

type createInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"categoryId"`
	ImageURL    string  `json:"imageUrl"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var in createInput
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	p, err := h.s.Create(r.Context(), in.Name, in.Description, in.Price, in.CategoryID, in.ImageURL)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	utils.JSON(w, 201, "product created", p)
}
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	q := utils.ParseQuery(r.URL.Query())
	list, err := h.s.List(r.Context(), q)
	if err != nil {
		utils.Error(w, 500, "failed to fetch products")
		return
	}
	utils.JSON(w, 200, "products fetched", list)
}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	p, err := h.s.Get(r.Context(), id)
	if err != nil {
		utils.Error(w, 404, "product not found")
		return
	}
	utils.JSON(w, 200, "product fetched", p)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var in createInput
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	p, err := h.s.Update(r.Context(), r.PathValue("id"), in.Name, in.Description, in.Price, in.CategoryID, in.ImageURL)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	utils.JSON(w, 200, "product updated", p)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.s.Delete(r.Context(), r.PathValue("id")); err != nil {
		utils.Error(w, 500, "failed to delete product")
		return
	}
	utils.JSON(w, 200, "product deleted", nil)
}

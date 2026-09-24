package auth

import (
	"encoding/json"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/http"
)

type Handler struct{ service *Service }

func NewHandler(s *Service) *Handler { return &Handler{service: s} }
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var in RegisterInput
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	id, err := h.service.Register(r.Context(), in)
	if err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	utils.JSON(w, 201, "registration successful", map[string]string{"id": id})
}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var in LoginInput
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	tokens, err := h.service.Login(r.Context(), in)
	if err != nil {
		utils.Error(w, 401, err.Error())
		return
	}
	utils.JSON(w, 200, "login successful", tokens)
}
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var in struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	tokens, err := h.service.Refresh(r.Context(), in.RefreshToken)
	if err != nil {
		utils.Error(w, 401, err.Error())
		return
	}
	utils.JSON(w, 200, "token refreshed", tokens)
}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var in struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}
	_ = json.NewDecoder(r.Body).Decode(&in)
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	h.service.Logout(r.Context(), in.RefreshToken)
	utils.JSON(w, 200, "logout successful", nil)
}
func (h *Handler) Forgot(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email string `json:"email" validate:"required,email"`
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	_ = h.service.ForgotPassword(r.Context(), in.Email)
	utils.JSON(w, 200, "if the email exists, a reset code was sent", nil)
}
func (h *Handler) Reset(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email    string `json:"email" validate:"required,email"`
		Code     string `json:"code" validate:"required,len=6,numeric"`
		Password string `json:"password" validate:"required,min=6,max=72"`
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		utils.Error(w, 400, "invalid body")
		return
	}
	if err := utils.ValidateStruct(in); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	if err := h.service.ResetPassword(r.Context(), in.Email, in.Code, in.Password); err != nil {
		utils.Error(w, 400, err.Error())
		return
	}
	utils.JSON(w, 200, "password reset successful", nil)
}

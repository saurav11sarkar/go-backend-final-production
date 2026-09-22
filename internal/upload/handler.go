package upload

import (
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/http"
	"strings"
)

type Handler struct {
	s     *Service
	maxMB int
}

func NewHandler(s *Service, maxMB int) *Handler { return &Handler{s: s, maxMB: maxMB} }
func (h *Handler) Image(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(int64(h.maxMB) * 1024 * 1024); err != nil {
		utils.Error(w, 400, "file is too large or invalid")
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		utils.Error(w, 400, "image is required")
		return
	}
	defer file.Close()
	content := header.Header.Get("Content-Type")
	if !strings.HasPrefix(content, "image/") {
		utils.Error(w, 400, "only image files are allowed")
		return
	}
	url, err := h.s.Image(r.Context(), file, header.Filename)
	if err != nil {
		utils.Error(w, 500, "image upload failed")
		return
	}
	utils.JSON(w, 200, "image uploaded", map[string]string{"url": url})
}

package upload

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
)

type Service struct{ cfg config.Config }

func New(cfg config.Config) (*Service, error) {
	if cfg.CloudName == "" || cfg.CloudAPIKey == "" || cfg.CloudAPISecret == "" {
		return nil, fmt.Errorf("Cloudinary is not configured")
	}
	return &Service{cfg: cfg}, nil
}

func (s *Service) Image(ctx context.Context, file multipart.File, filename string) (string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	folder := "go-backend/images"
	toSign := "folder=" + folder + "&timestamp=" + timestamp
	sum := sha1.Sum([]byte(toSign + s.cfg.CloudAPISecret))
	signature := hex.EncodeToString(sum[:])

	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	_ = mw.WriteField("api_key", s.cfg.CloudAPIKey)
	_ = mw.WriteField("timestamp", timestamp)
	_ = mw.WriteField("folder", folder)
	_ = mw.WriteField("signature", signature)
	part, err := mw.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}
	if _, err = io.Copy(part, file); err != nil {
		return "", err
	}
	if err = mw.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.cloudinary.com/v1_1/"+s.cfg.CloudName+"/image/upload", &body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return "", fmt.Errorf("cloudinary returned status %s", res.Status)
	}

	var out struct {
		SecureURL string `json:"secure_url"`
	}
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.SecureURL, nil
}

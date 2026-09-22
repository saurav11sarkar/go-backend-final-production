package link

import (
	"context"
	"errors"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"net/url"
	"strings"
)

type Service struct{ repo *Repository }

func NewService(repo *Repository) *Service { return &Service{repo: repo} }
func (s *Service) Create(ctx context.Context, title, raw string) (Link, error) {
	title = strings.TrimSpace(title)
	raw = strings.TrimSpace(raw)
	if title == "" {
		return Link{}, errors.New("title is required")
	}
	u, e := url.Parse(raw)
	if e != nil || u.Scheme == "" || u.Host == "" {
		return Link{}, errors.New("valid url is required")
	}
	l := Link{ID: utils.NewID(), Title: title, URL: raw}
	if e = s.repo.Create(ctx, l); e != nil {
		return Link{}, e
	}
	return l, nil
}
func (s *Service) List(ctx context.Context) ([]Link, error) { return s.repo.List(ctx) }

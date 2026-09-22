package category

import (
	"context"
	"errors"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"strings"
)

type Service struct{ repo *Repository }

func NewService(repo *Repository) *Service { return &Service{repo: repo} }
func (s *Service) Create(ctx context.Context, name string) (Category, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Category{}, errors.New("name is required")
	}
	c := Category{ID: utils.NewID(), Name: name, Slug: strings.ToLower(strings.ReplaceAll(name, " ", "-"))}
	if e := s.repo.Create(ctx, c); e != nil {
		return Category{}, e
	}
	return c, nil
}
func (s *Service) List(ctx context.Context, q utils.Query) ([]Category, error) {
	return s.repo.List(ctx, q)
}

func (s *Service) Update(ctx context.Context, id, name string) (Category, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Category{}, errors.New("name is required")
	}
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	if err := s.repo.Update(ctx, id, name, slug); err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Slug: slug}, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

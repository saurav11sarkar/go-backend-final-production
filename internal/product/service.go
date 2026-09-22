package product

import (
	"context"
	"errors"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"strings"
)

type Service struct{ repo *Repository }

func NewService(repo *Repository) *Service { return &Service{repo: repo} }
func (s *Service) Create(ctx context.Context, name, desc string, price float64, cat, image string) (Product, error) {
	name = strings.TrimSpace(name)
	if name == "" || price < 0 {
		return Product{}, errors.New("name and valid price are required")
	}
	p := Product{ID: utils.NewID(), Name: name, Description: desc, Price: price, CategoryID: cat, ImageURL: image}
	if e := s.repo.Create(ctx, p); e != nil {
		return Product{}, e
	}
	return s.Get(ctx, p.ID)
}
func (s *Service) Get(ctx context.Context, id string) (Product, error) { return s.repo.Get(ctx, id) }
func (s *Service) List(ctx context.Context, q utils.Query) ([]Product, error) {
	return s.repo.List(ctx, q)
}

func (s *Service) Update(ctx context.Context, id, name, desc string, price float64, cat, image string) (Product, error) {
	name = strings.TrimSpace(name)
	if name == "" || price < 0 {
		return Product{}, errors.New("name and valid price are required")
	}
	p := Product{ID: id, Name: name, Description: desc, Price: price, CategoryID: cat, ImageURL: image}
	if err := s.repo.Update(ctx, p); err != nil {
		return Product{}, err
	}
	return s.Get(ctx, id)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

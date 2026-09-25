package user

import (
	"context"
	"errors"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"strings"
)

type Service struct{ repo *Repository }

func NewService(repo *Repository) *Service                         { return &Service{repo: repo} }
func (s *Service) Me(ctx context.Context, id string) (User, error) { return s.repo.Me(ctx, id) }
func (s *Service) List(ctx context.Context, q utils.Query) ([]User, int, error) {
	return s.repo.List(ctx, q)
}
func (s *Service) UpdateProfile(ctx context.Context, id, name string) (User, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return User{}, errors.New("fullName is requireds")
	}
	if err := s.repo.UpdateProfile(ctx, id, name); err != nil {
		return User{}, err
	}
	return s.Me(ctx, id)
}
func (s *Service) SetProfilePicture(ctx context.Context, id, url string) (User, error) {
	if err := s.repo.SetProfilePicture(ctx, id, url); err != nil {
		return User{}, err
	}
	return s.Me(ctx, id)
}

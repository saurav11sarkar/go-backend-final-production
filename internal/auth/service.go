package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/email"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"strings"
	"time"
)

type Service struct {
	repo *Repository
	cfg  config.Config
	mail *email.Service
}

func NewService(repo *Repository, cfg config.Config, mail *email.Service) *Service {
	return &Service{repo: repo, cfg: cfg, mail: mail}
}

type RegisterInput struct {
	FullName string `json:"fullName" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (s *Service) Register(ctx context.Context, in RegisterInput) (string, error) {
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))
	if in.FullName == "" || in.Email == "" || len(in.Password) < 6 {
		return "", errors.New("fullName, valid email and password (min 6) are required")
	}
	hash, err := utils.HashPassword(in.Password)
	if err != nil {
		return "", err
	}
	id := utils.NewID()
	err = s.repo.CreateUser(ctx, id, in.FullName, in.Email, hash)
	if err != nil {
		return "", fmt.Errorf("email may already exist")
	}
	return id, nil
}
func (s *Service) Login(ctx context.Context, in LoginInput) (TokenResponse, error) {
	id, hash, role, status, err := s.repo.FindUserForLogin(ctx, strings.ToLower(strings.TrimSpace(in.Email)))
	if err != nil || !utils.CheckPassword(hash, in.Password) || status != "active" {
		return TokenResponse{}, errors.New("invalid credentials")
	}
	return s.tokens(ctx, id, role)
}
func (s *Service) tokens(ctx context.Context, id, role string) (TokenResponse, error) {
	access, err := utils.CreateToken(id, role, "access", s.cfg.AccessSecret, time.Duration(s.cfg.AccessMinutes)*time.Minute)
	if err != nil {
		return TokenResponse{}, err
	}
	refresh, err := utils.CreateToken(id, role, "refresh", s.cfg.RefreshSecret, time.Duration(s.cfg.RefreshDays)*24*time.Hour)
	if err != nil {
		return TokenResponse{}, err
	}
	hash, err := utils.HashPassword(refresh)
	if err != nil {
		return TokenResponse{}, err
	}
	expiresAt := time.Now().Add(time.Duration(s.cfg.RefreshDays) * 24 * time.Hour)
	if err := s.repo.SaveRefreshToken(ctx, utils.NewID(), id, hash, expiresAt); err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{access, refresh}, nil
}
func (s *Service) Refresh(ctx context.Context, token string) (TokenResponse, error) {
	claims, err := utils.ParseToken(token, s.cfg.RefreshSecret, "refresh")
	if err != nil {
		return TokenResponse{}, errors.New("invalid refresh token")
	}
	records, err := s.repo.FindRefreshTokens(ctx, claims.UserID)
	if err != nil {
		return TokenResponse{}, err
	}
	var tokenID string
	for _, record := range records {
		if utils.CheckPassword(record.TokenHash, token) {
			tokenID = record.ID
			break
		}
	}
	if tokenID == "" {
		return TokenResponse{}, errors.New("refresh token not found")
	}
	if err := s.repo.DeleteRefreshToken(ctx, tokenID); err != nil {
		return TokenResponse{}, err
	}
	return s.tokens(ctx, claims.UserID, claims.Role)
}

func (s *Service) Logout(ctx context.Context, token string) {
	claims, err := utils.ParseToken(token, s.cfg.RefreshSecret, "refresh")
	if err == nil {
		_ = s.repo.DeleteUserRefreshTokens(ctx, claims.UserID)
	}
}
func (s *Service) ForgotPassword(ctx context.Context, emailAddress string) error {
	id, err := s.repo.FindUserIDByEmail(ctx, strings.ToLower(strings.TrimSpace(emailAddress)))
	if err != nil {
		return nil
	}
	code := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
	err = s.repo.SaveResetCode(ctx, id, code, time.Now().Add(10*time.Minute))
	if err == nil {
		s.mail.Send(emailAddress, "Password Reset Code", "Your password reset code is: "+code)
	}
	return nil
}
func (s *Service) ResetPassword(ctx context.Context, emailAddress, code, password string) error {
	id, err := s.repo.FindUserByResetCode(ctx, strings.ToLower(strings.TrimSpace(emailAddress)), code)
	if err != nil {
		return errors.New("invalid or expired reset code")
	}
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return s.repo.ResetPassword(ctx, id, hash)
}

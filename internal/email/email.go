package email

import (
	"fmt"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/config"
	"net/smtp"
)

type Service struct{ cfg config.Config }

func New(cfg config.Config) *Service { return &Service{cfg: cfg} }
func (s *Service) Send(to, subject, body string) error {
	if s.cfg.SMTPHost == "" || s.cfg.SMTPUsername == "" {
		return fmt.Errorf("SMTP is not configured")
	}
	auth := smtp.PlainAuth("", s.cfg.SMTPUsername, s.cfg.SMTPPassword, s.cfg.SMTPHost)
	msg := []byte("From: " + s.cfg.SMTPFrom + "\r\nTo: " + to + "\r\nSubject: " + subject + "\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n" + body)
	return smtp.SendMail(fmt.Sprintf("%s:%d", s.cfg.SMTPHost, s.cfg.SMTPPort), auth, s.cfg.SMTPFrom, []string{to}, msg)
}

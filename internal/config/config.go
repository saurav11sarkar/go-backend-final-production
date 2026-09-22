package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppEnv         string
	Port           string
	DatabaseURL    string
	AccessSecret   string
	RefreshSecret  string
	AccessMinutes  int
	RefreshDays    int
	CorsOrigin     string
	CloudName      string
	CloudAPIKey    string
	CloudAPISecret string
	SMTPHost       string
	SMTPPort       int
	SMTPUsername   string
	SMTPPassword   string
	SMTPFrom       string
	MaxUploadMB    int
}

func Load() (Config, error) {
	port := getenv("PORT", "5000")
	accessMinutes, err := atoi(getenv("ACCESS_TOKEN_MINUTES", "15"))
	if err != nil {
		return Config{}, err
	}
	refreshDays, err := atoi(getenv("REFRESH_TOKEN_DAYS", "7"))
	if err != nil {
		return Config{}, err
	}
	smtpPort, err := atoi(getenv("SMTP_PORT", "587"))
	if err != nil {
		return Config{}, err
	}
	maxUpload, err := atoi(getenv("MAX_UPLOAD_MB", "5"))
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		AppEnv: getenv("APP_ENV", "development"), Port: port, DatabaseURL: os.Getenv("DATABASE_URL"),
		AccessSecret: os.Getenv("JWT_ACCESS_SECRET"), RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
		AccessMinutes: accessMinutes, RefreshDays: refreshDays, CorsOrigin: getenv("CORS_ORIGIN", "*"),
		CloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"), CloudAPIKey: os.Getenv("CLOUDINARY_API_KEY"), CloudAPISecret: os.Getenv("CLOUDINARY_API_SECRET"),
		SMTPHost: os.Getenv("SMTP_HOST"), SMTPPort: smtpPort, SMTPUsername: os.Getenv("SMTP_USERNAME"), SMTPPassword: os.Getenv("SMTP_PASSWORD"), SMTPFrom: os.Getenv("SMTP_FROM"),
		MaxUploadMB: maxUpload,
	}
	if cfg.DatabaseURL == "" || cfg.AccessSecret == "" || cfg.RefreshSecret == "" {
		return Config{}, fmt.Errorf("DATABASE_URL, JWT_ACCESS_SECRET and JWT_REFRESH_SECRET are required")
	}
	return cfg, nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
func atoi(v string) (int, error) {
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", v, err)
	}
	return n, nil
}

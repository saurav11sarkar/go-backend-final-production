package auth

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct{ db *pgxpool.Pool }

func NewRepository(db *pgxpool.Pool) *Repository { return &Repository{db: db} }

func (r *Repository) CreateUser(ctx context.Context, id, fullName, email, password string) error {
	_, err := r.db.Exec(ctx, `INSERT INTO users(id,full_name,email,password,role,status) VALUES($1,$2,$3,$4,'user','active')`, id, fullName, email, password)
	return err
}

func (r *Repository) FindUserForLogin(ctx context.Context, email string) (id, hash, role, status string, err error) {
	err = r.db.QueryRow(ctx, `SELECT id,password,role,status FROM users WHERE email=$1`, email).Scan(&id, &hash, &role, &status)
	return
}

func (r *Repository) SaveRefreshToken(ctx context.Context, id, userID, tokenHash string, expiresAt time.Time) error {
	_, err := r.db.Exec(ctx, `INSERT INTO refresh_tokens(id,user_id,token_hash,expires_at) VALUES($1,$2,$3,$4)`, id, userID, tokenHash, expiresAt)
	return err
}

func (r *Repository) FindRefreshTokens(ctx context.Context, userID string) ([]RefreshTokenRecord, error) {
	rows, err := r.db.Query(ctx, `SELECT id,token_hash FROM refresh_tokens WHERE user_id=$1 AND expires_at>NOW()`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []RefreshTokenRecord
	for rows.Next() {
		var item RefreshTokenRecord
		if err := rows.Scan(&item.ID, &item.TokenHash); err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, rows.Err()
}

func (r *Repository) DeleteRefreshToken(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM refresh_tokens WHERE id=$1`, id)
	return err
}

func (r *Repository) DeleteUserRefreshTokens(ctx context.Context, userID string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM refresh_tokens WHERE user_id=$1`, userID)
	return err
}

func (r *Repository) FindUserIDByEmail(ctx context.Context, email string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx, `SELECT id FROM users WHERE email=$1`, email).Scan(&id)
	return id, err
}

func (r *Repository) SaveResetCode(ctx context.Context, userID, code string, expiresAt time.Time) error {
	_, err := r.db.Exec(ctx, `UPDATE users SET reset_code=$1,reset_code_expires_at=$2,updated_at=NOW() WHERE id=$3`, code, expiresAt, userID)
	return err
}

func (r *Repository) ResetPassword(ctx context.Context, userID, password string) error {
	_, err := r.db.Exec(ctx, `UPDATE users SET password=$1,reset_code=NULL,reset_code_expires_at=NULL,updated_at=NOW() WHERE id=$2`, password, userID)
	return err
}

func (r *Repository) FindUserByResetCode(ctx context.Context, email, code string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx, `SELECT id FROM users WHERE email=$1 AND reset_code=$2 AND reset_code_expires_at>NOW()`, email, code).Scan(&id)
	return id, err
}

type RefreshTokenRecord struct {
	ID        string
	TokenHash string
}

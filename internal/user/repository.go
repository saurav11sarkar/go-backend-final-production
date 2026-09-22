package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"strings"
)

type Repository struct{ db *pgxpool.Pool }

func NewRepository(db *pgxpool.Pool) *Repository { return &Repository{db: db} }
func (r *Repository) Me(ctx context.Context, id string) (User, error) {
	var u User
	e := r.db.QueryRow(ctx, `SELECT id,full_name,email,role,status,COALESCE(profile_picture,''),created_at FROM users WHERE id=$1`, id).Scan(&u.ID, &u.FullName, &u.Email, &u.Role, &u.Status, &u.ProfilePicture, &u.CreatedAt)
	return u, e
}
func (r *Repository) List(ctx context.Context, q utils.Query) ([]User, int, error) {
	allowed := map[string]bool{"created_at": true, "full_name": true, "email": true}
	if !allowed[q.SortBy] {
		q.SortBy = "created_at"
	}
	var total int
	if e := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM users WHERE ($1='' OR full_name ILIKE '%'||$1||'%' OR email ILIKE '%'||$1||'%') AND ($2='' OR role=$2) AND ($3='' OR status=$3)`, q.Search, q.Role, q.Status).Scan(&total); e != nil {
		return nil, 0, e
	}
	rows, e := r.db.Query(ctx, `SELECT id,full_name,email,role,status,COALESCE(profile_picture,''),created_at FROM users WHERE ($1='' OR full_name ILIKE '%'||$1||'%' OR email ILIKE '%'||$1||'%') AND ($2='' OR role=$2) AND ($3='' OR status=$3) ORDER BY `+q.SortBy+` `+strings.ToUpper(q.SortOrder)+` LIMIT $4 OFFSET $5`, q.Search, q.Role, q.Status, q.Limit, q.Offset())
	if e != nil {
		return nil, 0, e
	}
	defer rows.Close()
	out := []User{}
	for rows.Next() {
		var u User
		if e := rows.Scan(&u.ID, &u.FullName, &u.Email, &u.Role, &u.Status, &u.ProfilePicture, &u.CreatedAt); e != nil {
			return nil, 0, e
		}
		out = append(out, u)
	}
	return out, total, rows.Err()
}
func (r *Repository) UpdateProfile(ctx context.Context, id, name string) error {
	_, e := r.db.Exec(ctx, `UPDATE users SET full_name=$1,updated_at=NOW() WHERE id=$2`, name, id)
	return e
}
func (r *Repository) SetProfilePicture(ctx context.Context, id, url string) error {
	_, e := r.db.Exec(ctx, `UPDATE users SET profile_picture=$1,updated_at=NOW() WHERE id=$2`, url, id)
	return e
}

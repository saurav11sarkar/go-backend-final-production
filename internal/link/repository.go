package link

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct{ db *pgxpool.Pool }

func NewRepository(db *pgxpool.Pool) *Repository { return &Repository{db: db} }
func (r *Repository) Create(ctx context.Context, l Link) error {
	_, e := r.db.Exec(ctx, `INSERT INTO links(id,title,url) VALUES($1,$2,$3)`, l.ID, l.Title, l.URL)
	return e
}
func (r *Repository) List(ctx context.Context) ([]Link, error) {
	rows, e := r.db.Query(ctx, `SELECT id,title,url,created_at::text FROM links ORDER BY created_at DESC`)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []Link{}
	for rows.Next() {
		var l Link
		if e := rows.Scan(&l.ID, &l.Title, &l.URL, &l.CreatedAt); e != nil {
			return nil, e
		}
		out = append(out, l)
	}
	return out, rows.Err()
}

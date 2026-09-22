package category

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"strings"
)

type Repository struct{ db *pgxpool.Pool }

func NewRepository(db *pgxpool.Pool) *Repository { return &Repository{db: db} }
func (r *Repository) Create(ctx context.Context, c Category) error {
	_, e := r.db.Exec(ctx, `INSERT INTO categories(id,name,slug) VALUES($1,$2,$3)`, c.ID, c.Name, c.Slug)
	return e
}
func (r *Repository) List(ctx context.Context, q utils.Query) ([]Category, error) {
	sort := "created_at"
	if q.SortBy == "name" {
		sort = "name"
	}
	rows, e := r.db.Query(ctx, `SELECT id,name,slug FROM categories WHERE ($1='' OR name ILIKE '%'||$1||'%') ORDER BY `+sort+` `+strings.ToUpper(q.SortOrder)+` LIMIT $2 OFFSET $3`, q.Search, q.Limit, q.Offset())
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []Category{}
	for rows.Next() {
		var c Category
		if e := rows.Scan(&c.ID, &c.Name, &c.Slug); e != nil {
			return nil, e
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

func (r *Repository) Update(ctx context.Context, id, name, slug string) error {
	_, err := r.db.Exec(ctx, `UPDATE categories SET name=$1,slug=$2,updated_at=NOW() WHERE id=$3`, name, slug, id)
	return err
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM categories WHERE id=$1`, id)
	return err
}

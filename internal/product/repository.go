package product

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/saurav11sarkar/go-backend-boilerplate/internal/utils"
	"strings"
)

type Repository struct{ db *pgxpool.Pool }

func NewRepository(db *pgxpool.Pool) *Repository { return &Repository{db: db} }
func (r *Repository) Create(ctx context.Context, p Product) error {
	_, e := r.db.Exec(ctx, `INSERT INTO products(id,name,description,price,image_url,category_id) VALUES($1,$2,$3,$4,$5,$6)`, p.ID, p.Name, p.Description, p.Price, p.ImageURL, p.CategoryID)
	return e
}
func (r *Repository) Get(ctx context.Context, id string) (Product, error) {
	var p Product
	e := r.db.QueryRow(ctx, `SELECT id,name,description,price,COALESCE(image_url,''),category_id,created_at::text FROM products WHERE id=$1`, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL, &p.CategoryID, &p.CreatedAt)
	return p, e
}
func (r *Repository) List(ctx context.Context, q utils.Query) ([]Product, error) {
	sort := "created_at"
	if q.SortBy == "name" {
		sort = "name"
	}
	rows, e := r.db.Query(ctx, `SELECT id,name,description,price,COALESCE(image_url,''),category_id,created_at::text FROM products WHERE ($1='' OR name ILIKE '%'||$1||'%') AND ($2='' OR category_id=$2) ORDER BY `+sort+` `+strings.ToUpper(q.SortOrder)+` LIMIT $3 OFFSET $4`, q.Search, q.CategoryID, q.Limit, q.Offset())
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []Product{}
	for rows.Next() {
		var p Product
		if e := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.ImageURL, &p.CategoryID, &p.CreatedAt); e != nil {
			return nil, e
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func (r *Repository) Update(ctx context.Context, p Product) error {
	_, err := r.db.Exec(ctx, `UPDATE products SET name=$1,description=$2,price=$3,image_url=$4,category_id=$5,updated_at=NOW() WHERE id=$6`, p.Name, p.Description, p.Price, p.ImageURL, p.CategoryID, p.ID)
	return err
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM products WHERE id=$1`, id)
	return err
}

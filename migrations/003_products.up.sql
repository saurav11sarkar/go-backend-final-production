CREATE TABLE IF NOT EXISTS products (
 id UUID PRIMARY KEY,
 name VARCHAR(200) NOT NULL,
 description TEXT NOT NULL DEFAULT '',
 price NUMERIC(12,2) NOT NULL CHECK (price >= 0),
 image_url TEXT,
 category_id UUID NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
 created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
 updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_products_category ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_products_created_at ON products(created_at);

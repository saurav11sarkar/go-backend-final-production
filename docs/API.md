# API Quick Reference

Base URL: `/api/v1`

## Auth

- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/refresh`
- `POST /auth/logout`
- `POST /auth/forgot-password`
- `POST /auth/reset-password`

## Users

- `GET /users` — admin, search/filter/pagination/sorting
- `GET /users/me` — authenticated user
- `PATCH /users/me` — authenticated user

## Categories

- `GET /categories` — search/pagination/sorting
- `POST /categories` — admin
- `PATCH /categories/{id}` — admin
- `DELETE /categories/{id}` — admin

## Products

- `GET /products` — search/filter/pagination/sorting
- `GET /products/{id}`
- `POST /products` — admin
- `PATCH /products/{id}` — admin
- `DELETE /products/{id}` — admin

Example:

```text
GET /api/v1/products?page=1&limit=10&search=phone&categoryId=<uuid>&sortBy=name&sortOrder=asc
```

## Upload

- `POST /uploads/image` — authenticated multipart field: `image`

## Health

- `GET /health`

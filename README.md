# Go Backend Boilerplate — Beginner Friendly + Production Ready

A modular Go backend for developers moving from Express.js / NestJS to Go.

**HTTP:** standard library `net/http`  
**Database:** PostgreSQL + `pgx`  
**ORM:** none  
**Router framework:** none  

The goal is simple: learn Go fundamentals while keeping a structure that can grow from small APIs to medium and larger production services.

## Architecture

```text
Request
  ↓
Global Middleware
  ↓
Module Router
  ↓
Auth / Role Middleware
  ↓
Handler
  ↓
Service
  ↓
Repository
  ↓
PostgreSQL
```

## Module structure

```text
internal/
├── auth/
│   ├── repository.go
│   ├── service.go
│   ├── handler.go
│   └── router.go
├── user/
│   ├── model.go
│   ├── repository.go
│   ├── service.go
│   ├── handler.go
│   └── router.go
├── product/
│   ├── model.go
│   ├── repository.go
│   ├── service.go
│   ├── handler.go
│   └── router.go
├── category/
│   ├── model.go
│   ├── repository.go
│   ├── service.go
│   ├── handler.go
│   └── router.go
├── link/
│   ├── model.go
│   ├── repository.go
│   ├── service.go
│   ├── handler.go
│   └── router.go
├── upload/
│   ├── cloudinary.go
│   ├── handler.go
│   └── router.go
├── email/
├── middleware/
├── database/
├── config/
├── routes/
└── utils/
```

For a new feature, repeat the same five-file module pattern. Example: `booking/model.go`, `booking/repository.go`, `booking/service.go`, `booking/handler.go`, `booking/router.go`.

## Included

### HTTP / Core
- `net/http` + `http.ServeMux` method/path routing
- Context propagation
- Request ID
- CORS
- Request logging
- Panic recovery
- Graceful shutdown
- Server timeouts

### Database
- PostgreSQL + `pgxpool`
- No GORM
- No ORM abstraction
- Connection pool configuration
- Separate `up/down` migrations for users, categories, products, links and refresh tokens

### Authentication
- Register / Login / Logout
- JWT access token
- JWT refresh token
- Refresh token rotation
- Refresh token hashing in database
- Auth middleware
- Role middleware (`user`, `admin`)
- Forgot password / reset password
- bcrypt password hashing

### API features
- CRUD pattern
- Search
- Filter
- Pagination
- Sorting
- Safe sort-column allowlists
- Consistent JSON responses
- Central application error helpers

### Files / External services
- Cloudinary image upload
- Multipart validation
- Upload size limit
- SMTP email

### Developer experience
- `.env.example`
- Makefile
- Air hot reload
- Dockerfile
- Docker Compose with PostgreSQL
- GitHub Actions CI
- API quick reference
- Project flow guide
- Error handling guide
- Express/Nest → Go learning path

## Node/Nest mental model

| Express / NestJS | Go |
|---|---|
| Controller | Handler |
| Service | Service |
| Repository / Prisma / Mongoose | Repository + SQL |
| Guard | Auth / Role Middleware |
| DTO | Struct |
| Module | Feature folder |
| Router | `router.go` + `http.ServeMux` |
| ConfigService | `config` |
| Exception Filter | `utils/error.go` + error response helper |
| `main.ts` | `cmd/server/main.go` |

## Run locally

```bash
cp .env.example .env
make tidy
make migrate-up
make dev
```

For tools:

```bash
make install-tools
```

## Example query

```text
GET /api/v1/products?page=1&limit=10&search=phone&categoryId=<uuid>&sortBy=name&sortOrder=asc
```

## Learning order

Start with only the `user` module:

```text
router → handler → service → repository → SQL
```

Then learn auth, product CRUD, query features, Cloudinary, email, migrations, Docker and deployment.

Do not add Redis, queues, WebSockets, gRPC or microservices until the core module flow feels comfortable. Add them later without changing the basic mental model.

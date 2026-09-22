# Go Learning Path From Express/NestJS

1. `cmd/server/main.go` — server startup and graceful shutdown
2. `internal/routes` — how `net/http` routes requests
3. `internal/user/router.go` — module routes
4. `internal/user/handler.go` — request/response handling
5. `internal/user/service.go` — business logic
6. `internal/user/repository.go` — SQL and database access
7. `internal/middleware` — middleware and JWT guard equivalent
8. `internal/auth` — access token + refresh token rotation
9. `internal/utils/query.go` — pagination, search, filter, sorting
10. `internal/upload` — Cloudinary multipart upload
11. `internal/email` — SMTP and password reset
12. `migrations` — schema changes per resource
13. Docker + Makefile + Air
14. Tests, Redis, queues, WebSocket/gRPC and microservices only after the core flow is comfortable

The goal is to move from a familiar NestJS/Express mental model to idiomatic Go gradually.

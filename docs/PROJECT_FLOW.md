# Project Flow

## Request lifecycle

```text
HTTP Request
    ↓
Global Middleware
    ├── CORS
    ├── Request ID
    ├── Logger
    └── Recovery
    ↓
Module Router
    ↓
Auth / Role Middleware (when required)
    ↓
Handler
    ↓
Service
    ↓
Repository
    ↓
PostgreSQL
    ↓
JSON Response
```

## Module rule

Every business feature follows the same pattern:

```text
internal/user/
├── model.go
├── repository.go
├── service.go
├── handler.go
└── router.go
```

Do the same for `product`, `category`, `order`, `booking`, `payment`, etc.

## Why this is beginner friendly

You can learn one module from top to bottom and then copy the same mental model to the next module. There is no Gin, no GORM, and no unnecessary enterprise abstraction layer.

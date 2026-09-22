# Error Handling

The boilerplate uses one simple rule:

```text
Repository/Service error
        ↓
Handler decides HTTP context
        ↓
utils.HandleError / utils.Error
        ↓
One JSON response shape
```

`internal/utils/error.go` contains the reusable `AppError` and `HandleError` functions.

For beginner-friendly handlers, `utils.Error(w, status, message)` is also available. It always returns:

```json
{
  "success": false,
  "message": "user not found",
  "error": {
    "code": "NOT_FOUND",
    "message": "user not found"
  }
}
```

The recovery middleware converts panics into the same JSON error style instead of returning plain-text `500` responses.

In larger projects, add more typed application errors in `error.go` instead of creating a new error package for every module.

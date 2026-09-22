FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:3.20
WORKDIR /app
RUN adduser -D appuser
COPY --from=builder /app/server ./server
COPY --from=builder /app/migrations ./migrations
USER appuser
EXPOSE 5000
CMD ["./server"]

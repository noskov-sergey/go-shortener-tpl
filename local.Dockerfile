FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/shortener/main.go

# Runtime stage
FROM alpine:latest

COPY migrations/ migrations/
COPY --from=builder /app/main main

ENTRYPOINT ["/main"]
# ===============================
# Stage 1: SIAPKAN BUILDER DAHULU
# ===============================
FROM golang:1.24-alpine AS builder

# Install git (kadang dibutuhkan go mod)
RUN apk add --no-cache git make
# Install migrate CLI + mysql driver
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

# Copy seluruh source
COPY . .

# Build binary
RUN go build -o main cmd/server/main.go



# ===============================
# Stage 2: Run
# ===============================
FROM alpine:latest

WORKDIR /app

# Install timezone & cert (best practice)
RUN apk add --no-cache ca-certificates tzdata bash make

# Copy binary
COPY --from=builder /app/main .

COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Copy migrations
COPY --from=builder /app/migrations ./migrations

# Copy Makefile (optional)
COPY Makefile .

# Optional: set timezone
ENV TZ=Asia/Jakarta

EXPOSE 3000

CMD ["./main"]

# FROM golang:1.24-alpine

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod tidy

# COPY . .

# RUN go build -o main cmd/server/main.go

# EXPOSE 3000

# CMD ["./main"]

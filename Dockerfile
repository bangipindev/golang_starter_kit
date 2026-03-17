# # ===============================
# # Stage 1: SIAPKAN BUILDER DAHULU
# # ===============================
# FROM golang:1.24-alpine AS builder

# WORKDIR /app

# # Install git (kadang dibutuhkan go mod)
# RUN apk add --no-cache git

# COPY go.mod go.sum ./
# RUN go mod tidy

# # Copy seluruh source
# COPY . .

# # Build binary
# RUN go build -o main cmd/server/main.go

# # ===============================
# # Stage 2: Run
# # ===============================
# FROM alpine:latest

# WORKDIR /app

# # Install timezone & cert (best practice)
# RUN apk add --no-cache ca-certificates tzdata

# # Copy binary
# COPY --from=builder /app/main .

# # Copy migrations
# # COPY --from=builder /app/migrations ./migrations

# # Optional: set timezone
# ENV TZ=Asia/Jakarta

# EXPOSE 3000

# CMD ["./main"]

FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main cmd/server/main.go

EXPOSE 3000

CMD ["./main"]

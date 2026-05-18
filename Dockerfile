FROM golang:1.25-alpine AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM deps AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/server ./cmd/main.go

FROM alpine:3.21
WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata netcat-openbsd
COPY --from=builder /bin/server /bin/server
COPY --from=builder /app/db/migrations ./db/migrations
COPY --from=builder /app/api/docs/swagger.json ./api/docs/swagger.json
COPY scripts/docker-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh"]

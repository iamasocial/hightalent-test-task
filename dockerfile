FROM golang:1.25.4-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app /app/app

COPY --from=builder /app/config/config.yaml /app/config/config.yaml

COPY --from=builder /app/migrations /app/migrations

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD wget -qO- http://localhost:8080/health || exit 1

CMD ["/app/app"]
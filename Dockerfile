# syntax=docker/dockerfile:1

FROM golang:1.24.5-alpine AS builder
WORKDIR /src
RUN apk add --no-cache build-base git

# download deps
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY . .

# build API
ENV CGO_ENABLED=0
RUN go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

# build SEED
RUN go build -trimpath -ldflags="-s -w" -o /out/seed ./cmd/seed

# runtime (distroless)
FROM gcr.io/distroless/static-debian12:latest AS runtime
WORKDIR /app
COPY --from=builder /out/api /app/api
COPY --from=builder /out/seed /app/seed
USER nonroot:nonroot
# Catatan: untuk service API, docker-compose akan menjalankan /app/api (default ENTRYPOINT/CMD di compose)
# Untuk service seed, compose sudah set entrypoint: ["/app/seed"]

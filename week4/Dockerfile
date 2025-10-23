FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache openssl

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux ldflags="-s -w" go build -o bin/mtls-server ./cmd/mtls && chmod +x entrypoint.sh

FROM alpine:latest

RUN apk add --no-cache curl

WORKDIR /app
COPY --from=builder /app/bin/ .
COPY --from=builder /app/entrypoint.sh .

ENTRYPOINT ["/app/entrypoint.sh"]


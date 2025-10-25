FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache openssl

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux ldflags="-s -w" go build -o bin/server .

FROM alpine:latest

RUN apk add --no-cache curl

WORKDIR /app
COPY --from=builder /app/bin/ .

CMD ["./app/server"]


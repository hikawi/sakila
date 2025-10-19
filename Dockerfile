FROM golang:1.25-alpine AS builder

RUN apk add --no-cache curl

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server ./

CMD [ "./server" ]


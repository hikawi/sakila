# Sakila

## Overview

A set of each week's assignments orchestrated via Docker Compose and Docker
Containers. I'll try to keep all weeks' tech stack similar but I can't promise
that.

Each week contains a separate `README.md` file to read about its setting up,
and the used technologies as a traceable markup.

## Installation

### Docker

Requirements:

- Docker Engine (using `containerd` or Docker Desktop)
- Docker Compose
- Docker Buildx

There's only one central database that is MariaDB (a MySQL dialect) that is
connected to by all other containers. This is managed by the `docker-compose.yml`
file at root.

You can use `--profile weekN` for each week's separated API.

```bash
# Example: Spin up week1
docker compose --profile week1 up
```

By default, there are these published ports:

- `3000` for Week 1.

### Native Golang

Requirements:

- Go v1.25 or later. (check with `go version`, here's mine `go version go1.25.3 darwin/arm64`)

For each project, you can either run it directly, or build the binary to run it instead.

```bash
# Download dependencies first
go mod download

# If you want to run it directly
cd week1
go run

# If you want to build the binary
cd week1
CGO_ENABLED=0 GOOS="your-os" GOARCH="your-arch" ldflags="-w -s" go build -o server
./server
```

If you want to go native, you need to provide the correct environment values. A
local `.env` file is supported.

## Week Overview

| Week | Contents        |
| :--: | --------------- |
|  1   | Actors REST API |

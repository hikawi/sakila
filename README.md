# Sakila

## Overview

A set of each week's assignments orchestrated via Docker Compose and Docker
Containers. I'll try to keep all weeks' tech stack similar but I can't promise
that.

Each week contains a separate `README.md` file to read about its setting up,
and the used technologies as a traceable markup.

## Development's Environment

Make sure you have setup `.env` file. If you use Compose, make sure that `.env`
file stays at the same directory as `docker-compose.yml`. If you want to run
native, `godotenv` also supports loading an `.env` at each week's root.

**I recommend cleaning and rebuilding the MySQL server through each run as
each week might modify the database.**

### Docker

Requirements:

- Docker Engine (using `containerd` or Docker Desktop)
- Docker Compose
- Docker Buildx

Docker Compose is setup to use multiple profiles to allow separate dependencies
for each week's assignment. For example, week 1 wouldn't need week 2's swagger
UI generator.

```bash
# Example: spin up week1
docker compose --profile week1 up
```

You may use `--build` to force a rebuild, but usually a developer's compose file
would not need it.

By default, there are these published ports:

- `3306` for MySQL.
- `3000` for Week 1.
- `3001` for Week 2.
- `3002` for Week 3.
- `3003`, `3004` and `3005` for Week 4.

I know usually when you isolate profiles, these will only map in the correct
ports for that specific week. But if in any case, you need to spin up every
service, this port difference is for that reason.

### Go Native

**This is better for development**.

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

All weeks are using port `80`. Except for `week4`.

## Production Environment

The difference between a Production environment and a Developer environment is
that I setup for the dev environment to support hot reloads, as you change files,
Air will re-compile and re-build and re-deploy the server.

A production environment has no such mechanisms, using these, if any changes
happen, you have to force a rebuild. This is meant to be placed for smoother CI/CD
integrations.

## Week Overview

| Week |                                                                                                                                                                                                       CI/CD Status                                                                                                                                                                                                       | Contents                                                   |
| :--: | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: | ---------------------------------------------------------- |
|  1   |                                                                                                                                               ![GitHub Actions Workflow Status for Week 1](https://img.shields.io/github/actions/workflow/status/hikawi/sakila/week1.yml)                                                                                                                                                | Actors REST API                                            |
|  2   |                                                                                                                                               ![GitHub Actions Workflow Status for Week 2](https://img.shields.io/github/actions/workflow/status/hikawi/sakila/week2.yml)                                                                                                                                                | Films REST API + Request Validator + Swagger Documentation |
|  3   |                                                                                                                                               ![GitHub Actions Workflow Status for Week 3](https://img.shields.io/github/actions/workflow/status/hikawi/sakila/week3.yml)                                                                                                                                                | Films REST API + ODF stack for logging and monitoring (1)  |
|  4   | ![GitHub Actions Workflow Status for Week 4 mTLS](https://img.shields.io/github/actions/workflow/status/hikawi/sakila/week4-mtls.yml) ![GitHub Actions Workflow Status for Week 4 API Token](https://img.shields.io/github/actions/workflow/status/hikawi/sakila/week4-apitoken.yml) ![GitHub Actions Workflow Status for Week 4 JWT](https://img.shields.io/github/actions/workflow/status/hikawi/sakila/week4-jwt.yml) | Actors REST API + API Security (2)                         |

Notes on Tech/Libraries used:

- (1): Includes **O**penSearch, OpenSearch **D**ashboards and **F**luentbit.
- (2): Includes mTLS with Go's `x509`, API Token and JWT.

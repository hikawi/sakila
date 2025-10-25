# Week 4

## Overview

This week's objective is API Security while also providing an Authentication
layer to identify the requester, through three stages.

Technology Stack employed:

- Golang 1.25
  - GIN Web Framework
  - GORM for ORM solution
  - Auth0's Go JWT Middleware

## Setting up

Required environment variables:

- `MYSQL_DSN`: The DSN for MySQL connection.
- `JWT_SECRET`: The secret for JWT.

### Docker (Recommended)

Refer to the root directory's `docker-compose.yml`.

```bash
docker compose --profile week4-jwt up
```

The server will run on port `3005` mapped to container's `80`.

### Native

I recommend [air](https://github.com/air-verse/air) as the runner as it supports
hot reloads. Otherwise, running the following command and killing the process as
you go also works:

```bash
go run .
```

## Solution

Server will run on port `80` if run directly.

**Situation**: Anyone can register arbitrarily at any time and you need to
identify all of them.

**Solution**: JSON Web Tokens.

Steps:

1. Generate a JWT everytime the user logs in, including several claims, subject.
2. The user sends it via an Authorization header or via cookies.
3. Logging out will blacklist the JWT.

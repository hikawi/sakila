# Week 4

## Overview

This week's objective is API Security while also providing an Authentication
layer to identify the requester, through three stages.

Technology Stack employed:

- Golang 1.25
  - GIN Web Framework
  - GORM for ORM solution

## Setting up

### Docker (Recommended)

Refer to the root directory's `docker-compose.yml`.

```bash
docker compose --profile week4-apitoken up
```

The server will run on port `3004` mapped to container's `80`.

### Native

I recommend [air](https://github.com/air-verse/air) as the runner as it supports
hot reloads. Otherwise, running the following command and killing the process as
you go also works:

```bash
go run ./cmd/server
```

## Solution

Server will run on port `80` if run directly.

**Situation**: A small number of clients can connect and can be identified by
the server.

**Solution**: API Tokens, similar to OpenAI token or GitHub PAT. Uniquely
identify a user using an opaque token.

Steps:

1. Generate a Root CA (Certificate Authority). Any certificates signed by this
   authority, the server will trust.
2. Generate a Server Key (just random mostly), then a CSR (Certificate Signing
   Request). It requests the CA to sign a Certificate using the CSR (which has
   fields like claims in JWT), producing a CRT file. The server will present
   both the CRT file and the KEY file.
3. For each client, you can generate a Client Key, using the same mechanisms.
4. After both are signed by the Root CA, the server can trust and identify the
   client by looking at the digital certificate presented.

### Small number of Clients

Server will run on port `80` if run directly.

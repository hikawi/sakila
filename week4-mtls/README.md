# Week 4 - MTLS

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
docker compose --profile week4-mtls up
```

The server will run on port `3003` mapped to container's `443`.

### Native

I recommend [air](https://github.com/air-verse/air) as the runner as it supports
hot reloads. Otherwise, running the following command and killing the process as
you go also works:

```bash
go run .
```

### Common Problems

#### Can't find CA certificate

You might see:

```plaintext
fatal: can't find root CA certificate
```

You need to mount a folder named `certs` into `/app/certs` into the container,
if running with Docker, or when running natively, you need a folder named `certs`
at the working directory.

The folder `certs` must contain the following:

- `ca.crt`: The Root CA certificate.
- `server.key`: The key to identify the server.
- `server.crt`: The certificate to identify the server, signed by `ca.crt`.

For example, if you need to build and run by yourself:

```bash
docker build -t week4
docker run --name week4 -p 3004:443 -v ./certs:/app/certs week4
```

## Solution

Server will run on port `443` if run directly.

**Situation**: Clients that are trusted inside a domain or an inner circle
can connect, for example, in a law firm or a medical setting. Under no
circumstances are uncontrolled machines connected.

**Solution**: mutual TLS. Simple TLS is that the server has to present a certificate
to prove that it is who it claims to be. Mutual TLS requires the client also to
present a certificate.

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

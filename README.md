### Project Description

Conio - A simple social network

### Tools

- Postgres driver - `go get github.com/jackc/pgx/v5`

- Branca Token (alternative to JWT) - `go get -u github.com/hako/branca`

### Directory Structure

    internal/
        service/ (Holds business logic to back the APIs)
        handler/ (HTTP transport layer (can be swapped for GRPC or graphql))

## Database

To start an insecure cockroach db - `cockroach start-single-node --insecure --host 127.0.0.1`
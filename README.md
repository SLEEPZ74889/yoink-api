> © 2026 Eric Dopsonder. All rights reserved.
> Unauthorized copying, distribution, or use of this source code is strictly prohibited.

# yoink 🪝

> Steal the length. Keep the data.

A fast, minimal URL shortener with real-time analytics — built with Go, Chi, and PostgreSQL.

Every click is tracked: browser, OS, device type, referrer, and a privacy-safe IP hash. No bloat, no framework magic. Just clean layered architecture and raw SQL.

---

## Endpoints

```
GET  /health       — health check
POST /links        — create a short link
GET  /{slug}       — redirect + track click
```

### Create a link

```bash
curl -X POST http://localhost:8080/links \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com", "slug": "myslug", "title": "Example"}'
```

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "slug": "myslug",
  "short_url": "https://yoink.link/myslug",
  "original_url": "https://example.com",
  "created_at": "2026-05-20T14:00:00Z"
}
```

`slug` and `title` are optional. If no slug is provided, a random 6-character one is generated.

---

## Stack

- **Go** — no framework, no magic
- **Chi** — lightweight HTTP router
- **PostgreSQL (Neon)** — serverless postgres
- **pgx/v5** — native Go postgres driver with connection pooling

## Architecture

```
handler  →  service  →  repository  →  database
```

- **handler** — reads HTTP requests, writes HTTP responses
- **service** — business logic, validation, slug generation
- **repository** — SQL queries, nothing else
- **model** — shared data structures

Layers communicate through interfaces, which keeps each layer testable in isolation.

---

## Getting started

```bash
git clone https://github.com/SLEEPZ74889/yoink-api
cd yoink-api

cp .env.example .env
# add your Postgres connection string to .env

make migrate   # run DB migrations
make run       # start server on :8080
```

## Development

```bash
make test    # run all tests
make build   # build binary → bin/yoink
make fmt     # format code
make tidy    # go mod tidy
```

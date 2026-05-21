> © 2026 Eric Dopsonder. All rights reserved.
> Unauthorized copying, distribution, or use of this source code is strictly prohibited.

# yoink 🪝

> Steal the length. Keep the data.

A fast, minimal URL shortener with real-time analytics — built with Go, Chi, and PostgreSQL.

---

## What it does

- Shortens any URL to a clean slug (e.g. `yoink.link/x7k2mq`)
- Tracks every click: device, browser, OS, country, city, referrer
- Custom slugs supported
- Links can expire or be deactivated

## Stack

- **Go** — fast, no bloat
- **Chi** — lightweight HTTP router
- **PostgreSQL (Neon)** — serverless, scales to zero
- **pgx/v5** — native Go Postgres driver

## Architecture

Clean layered architecture following enterprise patterns:

```
handler  →  service  →  repository  →  database
```

Each layer has one job. No god objects, no spaghetti.

## API

```bash
# Health check
GET /health

# Create a short link
POST /links
{
  "url": "https://example.com",
  "slug": "myslug",   # optional
  "title": "My Link"  # optional
}
```

## Getting started

```bash
# Clone & run
git clone https://github.com/SLEEPZ74889/yoink-api
cd yoink-api

# Add your Postgres connection string
echo 'DATABASE_URL=your_connection_string' > .env

# Run migrations
make migrate

# Start the server
make run
```

Server runs on `:8080`.

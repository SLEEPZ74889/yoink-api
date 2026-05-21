-- internal/db/migrations/001_init.sql

CREATE TABLE links (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug         VARCHAR(20) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    title        VARCHAR(255),
    active       BOOLEAN DEFAULT true,
    created_at   TIMESTAMP DEFAULT NOW(),
    expires_at   TIMESTAMP
);

CREATE TABLE clicks (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    link_id    UUID REFERENCES links(id) ON DELETE CASCADE,
    clicked_at TIMESTAMP DEFAULT NOW(),
    country    VARCHAR(2),
    city       VARCHAR(100),
    referrer   TEXT,
    browser    VARCHAR(50),
    os         VARCHAR(50),
    device     VARCHAR(20),
    ip_hash    VARCHAR(64)
);

CREATE INDEX idx_links_slug      ON links(slug);
CREATE INDEX idx_clicks_link_id  ON clicks(link_id);
CREATE INDEX idx_clicks_clicked_at ON clicks(clicked_at);
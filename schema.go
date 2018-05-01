// Package pgtime is for testing how PostgreSQL timestamptz is stored/retrieved
// using drivers such as pgx and pq.
package pgtime

import "time"

// User is an entity being stored in Postgres.
type User struct {
	ID        uint64
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserSchema is db schema to test Postgres driver time handling.
const UserSchema = `
CREATE TABLE IF NOT EXISTS account (
    id bigserial,
	username varchar(40) NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,

    PRIMARY KEY(id),
    UNIQUE(username)
);
`

// Queries to create and read user accounts.
const (
	CreateQ = "INSERT INTO account (username, created_at, updated_at) VALUES ($1, $2, $3)"
	ReadQ   = "SELECT id, created_at, updated_at FROM account WHERE username=$1"
)

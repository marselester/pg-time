package pgx_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	// pgx driver registers itself as being available to the database/sql package.
	_ "github.com/jackc/pgx/stdlib"

	"github.com/marselester/pg-time"
)

func TestPgx(t *testing.T) {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "account",
		User:     "account",
		Password: "swordfish",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	if _, err = conn.Exec("SET TIME ZONE 'UTC'"); err != nil {
		t.Fatal(err)
	}
	if _, err = conn.Exec(pgtime.UserSchema); err != nil {
		t.Fatal(err)
	}
	if _, err = conn.Exec("DELETE FROM account"); err != nil {
		t.Fatal(err)
	}

	bob := pgtime.User{
		Username:  "bob",
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UpdatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	if _, err = conn.Exec(pgtime.CreateQ, bob.Username, bob.CreatedAt, bob.UpdatedAt); err != nil {
		t.Fatal(err)
	}
	u := pgtime.User{}
	if err = conn.QueryRow(pgtime.ReadQ, "bob").Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt); err != nil {
		t.Fatal(err)
	}

	if !bob.UpdatedAt.Equal(u.UpdatedAt) {
		t.Errorf("wanted UpdatedAt %v got %v", bob.UpdatedAt, u.UpdatedAt)
	}
	if bob.CreatedAt != u.CreatedAt {
		t.Errorf("wanted CreatedAt %v got %v", bob.CreatedAt, u.CreatedAt)
	}
}

func TestPgxParseConnectionString(t *testing.T) {
	dataSourceName := "postgres://account:swordfish@localhost/account?timezone=UTC"
	config, err := pgx.ParseConnectionString(dataSourceName)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	if _, err = conn.Exec(pgtime.UserSchema); err != nil {
		t.Fatal(err)
	}
	if _, err = conn.Exec("DELETE FROM account"); err != nil {
		t.Fatal(err)
	}

	bob := pgtime.User{
		Username:  "bob",
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UpdatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	if _, err = conn.Exec(pgtime.CreateQ, bob.Username, bob.CreatedAt, bob.UpdatedAt); err != nil {
		t.Fatal(err)
	}
	u := pgtime.User{}
	if err = conn.QueryRow(pgtime.ReadQ, "bob").Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt); err != nil {
		t.Fatal(err)
	}

	if !bob.UpdatedAt.Equal(u.UpdatedAt) {
		t.Errorf("wanted UpdatedAt %v got %v", bob.UpdatedAt, u.UpdatedAt)
	}
	if bob.CreatedAt != u.CreatedAt {
		t.Errorf("wanted CreatedAt %v got %v", bob.CreatedAt, u.CreatedAt)
	}
}

func TestPgxDatabaseSQL(t *testing.T) {
	dataSourceName := "postgres://account:swordfish@localhost/account?timezone=UTC"
	pool, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	if _, err = pool.Exec(pgtime.UserSchema); err != nil {
		t.Fatal(err)
	}
	if _, err = pool.Exec("DELETE FROM account"); err != nil {
		t.Fatal(err)
	}

	bob := pgtime.User{
		Username:  "bob",
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UpdatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	if _, err = pool.Exec(pgtime.CreateQ, bob.Username, bob.CreatedAt, bob.UpdatedAt); err != nil {
		t.Fatal(err)
	}
	u := pgtime.User{}
	if err = pool.QueryRow(pgtime.ReadQ, "bob").Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt); err != nil {
		t.Fatal(err)
	}

	if !bob.UpdatedAt.Equal(u.UpdatedAt) {
		t.Errorf("wanted UpdatedAt %v got %v", bob.UpdatedAt, u.UpdatedAt)
	}
	if bob.CreatedAt != u.CreatedAt {
		t.Errorf("wanted CreatedAt %v got %v", bob.CreatedAt, u.CreatedAt)
	}
}

func TestPgxTimestamp(t *testing.T) {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "account",
		User:     "account",
		Password: "swordfish",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	if _, err = conn.Exec("SET TIME ZONE 'UTC'"); err != nil {
		t.Fatal(err)
	}
	if _, err = conn.Exec(pgtime.UserSchema); err != nil {
		t.Fatal(err)
	}
	if _, err = conn.Exec("DELETE FROM account"); err != nil {
		t.Fatal(err)
	}

	bob := pgtime.User{
		Username:  "bob",
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UpdatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	if _, err = conn.Exec(pgtime.CreateQ, bob.Username, bob.CreatedAt, bob.UpdatedAt); err != nil {
		t.Fatal(err)
	}

	var createdAt pgtype.Timestamptz
	var updatedAt pgtype.Timestamptz
	u := pgtime.User{}
	if err = conn.QueryRow(pgtime.ReadQ, "bob").Scan(&u.ID, &createdAt, &updatedAt); err != nil {
		t.Fatal(err)
	}
	u.CreatedAt = createdAt.Time
	u.UpdatedAt = updatedAt.Time

	if !bob.UpdatedAt.Equal(u.UpdatedAt) {
		t.Errorf("wanted UpdatedAt %v got %v", bob.UpdatedAt, u.UpdatedAt)
	}
	if bob.CreatedAt != u.CreatedAt {
		t.Errorf("wanted CreatedAt %v got %v", bob.CreatedAt, u.CreatedAt)
	}
}

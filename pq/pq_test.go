package pq_test

import (
	"database/sql"
	"testing"
	"time"

	// postgres driver registers itself as being available to the database/sql package.
	_ "github.com/lib/pq"

	"github.com/marselester/pg-time"
)

func TestPq(t *testing.T) {
	dataSourceName := "postgres://account:swordfish@localhost/account?sslmode=disable"
	pool, err := sql.Open("postgres", dataSourceName)
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

# PostgreSQL Timestamptz and Go Idiosyncrasies

PostgreSQL converts time from the connection's time zone to UTC on storage,
and from UTC to the connection's time zone on retrieval.
However I was surprised to find that [pgx PostgreSQL driver](https://github.com/jackc/pgx)
doesn't return `time.Time` in UTC format despite configuring connection with

- `postgres://account:swordfish@localhost/account?timezone=UTC`
- `SET TIME ZONE 'UTC'`
- `SET timezone TO 'UTC'`

For example, `2009-11-10 23:00:00 +0000 UTC` is retrieved from db as `2009-11-11 06:00:00 +0700 +07`.
Though the dates are equal, the UTC format is what I expected.

I compared the behavior with [pq PostgreSQL driver](https://github.com/lib/pq/) and it returns UTC date.

```sh
$ make docker_run_postgres
$ go test ./...
```

Another caveat is `time.Time` nanosecond resolution and timestamptz microsecond resolution.
The current solution is to use [Truncate](https://github.com/lib/pq/issues/227).

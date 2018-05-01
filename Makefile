TEST_PGPORT := 5432
TEST_PGDATABASE := account
TEST_PGUSER := account
TEST_PGPASSWORD := swordfish
export TEST_PGPORT TEST_PGDATABASE TEST_PGUSER TEST_PGPASSWORD

docker_run_postgres:
	docker run --rm -p 5432:5432 -e POSTGRES_USER=$(TEST_PGUSER) -e POSTGRES_PASSWORD=$(TEST_PGPASSWORD) postgres:10.3-alpine

test:
	go test ./...

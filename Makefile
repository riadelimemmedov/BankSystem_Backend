create_migration:
	migrate create -ext sql -dir db/migration -seq init_schema
migrate_up:
	migrate -path db/migration/ -database "postgresql://postgres:123321@localhost:6432/simple_bank?sslmode=disable" --verbose up
migrate_down:
	migrate -path db/migration/ -database "postgresql://postgres:123321@localhost:6432/simple_bank?sslmode=disable" --verbose down


sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go simplebank/db/sqlc Store

postgres:
	docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123321 -p 6432:5432 -d postgres:12-alpine
get_postgres:
	docker exec -it postgres psql -U postgres -d postgres
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
get_db:
	docker exec -it postgres psql --username=postgres --dbname=postgres simple_bank
dropdb:
	docker exec -it postgres dropdb -U postgres simple_bank

get_command_history:
	history | grep "docker run"

get_accounts:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM accounts"
get_transfers:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM transfers"
get_entries:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM entries"
get_users:
	docker exec -it postgres psql -U postgres -d simple_bank -c "SELECT * FROM users"
delete_accounts:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE accounts RESTART IDENTITY CASCADE"
delete_entries:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE entries RESTART IDENTITY CASCADE"
delete_transfers:
	docker exec -it postgres psql -U postgres -d simple_bank -c "TRUNCATE transfers RESTART IDENTITY CASCADE"
update_accounts:
	docker exec -it postgres psql -U postgres -d simple_bank -c "UPDATE accounts SET currency='EUR' WHERE ID = $(ACCOUNT_ID)"

run_test:
	go test -v -cover ./...
clean_test_cache:
	go clean -testcache

.PHONY: create_migration,migrate_up,migrate_down,migrate_sql
.PHONY: sqlc
.PHONY: mock
.PHONY: postgres,get_postgres,createdb,dropdb
.PHONY: get_command_history
.PHONY: get_accounts get_transfers get_entries delete_accounts delete_entries delete_transfers update_accounts
.PHONY: run_test,clean_test_cache
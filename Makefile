CURRENT_DIR=$(shell pwd)
POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/buy_event?sslmode=disable'

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

run-unit-tests:
	cd ${CURRENT_DIR}/storage/postgres && go test -v


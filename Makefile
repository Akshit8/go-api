postgres:
	docker-compose up -d

createdb:
	docker exec -it postgresdb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresdb dropdb simple_bank

migrationup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrationup migrationdown sqlc test
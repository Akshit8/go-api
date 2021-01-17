postgres:
	docker-compose up -d

createdb:
	docker exec -it postgresdb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresdb dropdb simple_bank

migrationup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrationup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migrationdown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrationdown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

git:
	git add .
	git commit -m "$(msg)"
	git push origin master

server:
	export APP_ENV=dev && go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/Akshit8/go-api/db/sqlc Store

.PHONY: postgres createdb dropdb migrationup migrationdown migrationup1 migrationdown1 sqlc test git server mock
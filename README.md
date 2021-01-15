# go-api

## DB Architecture
<img src=".github/assets/BankServiceSchema.png">

## Setting PostgresDB for development with docker
```bash
# make sure to run in root folder
docker-compose up -d

# getting a postgres shell inside postgres container with root user
docker exec -it postgresdb psql -U root
```

## Set databse migration
```bash
## using migrate cli

## create migration folder
db/migration

# genertate init migration
migrate create -ext sql -dir db/migration -seq init_schema

## write up the initial migrations usind BankServiceSchema.sql

# set bank service db
docker exec -it postgresdb createdb --username=root --owner=root simple_bank

# migrate up
migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## CRUD Golang code from SQL
```bash
## using sqlc

## db/sql vs GORM vs sqlx vs sqlc

# to generate sqlc.yaml
sqlc init

# write CRUD sql queries inside db/query
```

## Golang unit tests for Database CRUD
```bash
## the db/sql is just query client and doesn't features any driver
## so for this we need to install postgres driver for go
go get github.com/lib/pq

## for writing single liners assertions and prevent if-else logic, install
go get github.com/stretchr/testify

# run test coverage
go test -v -cover ./...
```

## DB transaction lock
```bash
docker exec -it postgresdb psql -U root -d simple_bank

# Begin a transaction in postgres
BEGIN;

# Rollback a transaction
ROLLBACK;

# Commit a transaction
COMMIT;

# BEGIN; -> T1
SELECT * FROM accounts WHERE id = 1 FOR UPDATE;

# BEGIN; -> T2
SELECT * FROM accounts WHERE id = 1 FOR UPDATE;
## will block the query untile T1 is commit

# T1 update
UPDATE accounts SET balance = 250 WHERE id = 1;

# Commit T1 to release the block on T2
COMMIT;
```

## Makefile specs
- **postgres** - setup postgress with compose
- **createdb** - create a service db inside postgres
- **dropdb** - remove servie db
- **migrationup** - migrate db to new migrations
- **migrationdown** - rollback db to previous stage
- **sqlc** - generate golang db functions from sql queries
- **test** - run tests in all packages and prints verbose with line coverage

## References
[dbSchemaHelper](https://dbdiagram.io) <br>
[migration-tool](https://github.com/golang-migrate/migrate)<br>
[crud-code-generator](https://github.com/kyleconroy/sqlc)<br>
[go-postgres-driver](https://github.com/lib/pq)<br>
[common-assertion-toolkit](https://github.com/stretchr/testify)<br>
[accessing-host-network-inside-deocker-container](https://stackoverflow.com/questions/24319662/from-inside-of-a-docker-container-how-do-i-connect-to-the-localhost-of-the-mach#:~:text=Use%20%2D%2Dnetwork%3D%22host%22,for%20Linux%2C%20per%20the%20documentation.)<br>

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License
# go-api

## DB Architecture
### v1
<img src=".github/assets/BankServiceSchemaV1.png">

### v2
<img src=".github/assets/BankServiceSchemaV2.png">

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

## DB transaction lock and handling deadlock in postgres
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

## Isolation Levels in mysql & postgres
```
read-uncommited
read-commited
repeatable-read
serializable
```

## Loading config and env vars with go viper
Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile or Java properties formats.
```bash
# setup config file locally
cat .sample.env > app.env

# populate the app.env file
```

## Using GoMock to create mock DB for testing HTTP Api
```bash
# installing mockgen
go get github.com/golang/mock/mockgen@v1.4.4

# checking binary status
which mockgen
mockgen --version

# mockgen has two modes of operation: source and reflect.
Source mode generates mock interfaces from a source file.
It is enabled by using the -source flag
Example:
    mockgen -source=foo.go [other options]

Reflect mode generates mock interfaces by building a program
that uses reflection to understand interfaces. It is enabled
by passing two non-flag arguments: an import path, and a
comma-separated list of symbols.
Example:
    mockgen database/sql/driver Conn,Driver

# generate some mock code
mockgen -destination db/mock/store.go -package mockdb github.com/Akshit8/go-api/db/sqlc Store
```

## Migrating DB to v2
```bash
migrate create -ext sql -dir db/migration -seq add_users

make migrationup # would throw error due to foreign key constrant

docker exec -it postgresdb psql -U root -d simple_bank
SELECT * FROM schema_migrations;
UPDATE schema_migrations SET dirty=false where version=2;

make migrationdown
```

## Switiching app stage automatically
```bash
export APP_ENV=dev && go run main.go
```

## Makefile specs
- **postgres** - setup postgress with compose
- **createdb** - create a service db inside postgres
- **dropdb** - remove servie db
- **migrationup** - migrate db to new migrations
- **migrationdown** - rollback db to previous stage
- **migrationup1** - migrates to 1 up
- **migrationdown1** - rollbacks to 1 down
- **sqlc** - generate golang db functions from sql queries
- **test** - run tests in all packages and prints verbose with line coverage
- **git** - git add - commit - push commands
- **server** - start REST server
- **mock** - generated mockdb stub for api testing

## References
[dbSchemaHelper](https://dbdiagram.io)<br>
[migration-tool](https://github.com/golang-migrate/migrate)<br>
[crud-code-generator](https://github.com/kyleconroy/sqlc)<br>
[go-postgres-driver](https://github.com/lib/pq)<br>
[common-assertion-toolkit](https://github.com/stretchr/testify)<br>
[accessing-host-network-inside-deocker-container](https://stackoverflow.com/questions/24319662/from-inside-of-a-docker-container-how-do-i-connect-to-the-localhost-of-the-mach#:~:text=Use%20%2D%2Dnetwork%3D%22host%22,for%20Linux%2C%20per%20the%20documentation.)<br>
[postgres-lock](https://wiki.postgresql.org/wiki/Lock_Monitoring)<br>
[handling-deadlock-go](https://www.youtube.com/watch?v=G2aggv_3Bbg&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=7)<br>
[github-action-postgres](https://docs.github.com/en/free-pro-team@latest/actions/guides/creating-postgresql-service-containers)<br>
[gin](https://github.com/gin-gonic/gin)<br>
[go-validator](https://godoc.org/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags)<br>
[go-viper](https://github.com/spf13/viper)<br>

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License
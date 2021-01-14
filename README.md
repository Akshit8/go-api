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


```

## Makefile specs
- **postgres** - setup postgress with compose
- **createdb** - create a service db inside postgres
- **dropdb** - remove servie db
- **migartionup** - migrate db to new migrations
- **migartiondown** - rollback db to previous stage

## References
[dbSchemaHelper](https://dbdiagram.io) <br>
[migration-tool](https://github.com/golang-migrate/migrate)<br>
[crud-code-generator](https://github.com/kyleconroy/sqlc)<br>

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License
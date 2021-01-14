# go-api

## DB Architecture
<img src=".github/assets/BankServiceSchema.png">

## References
[dbSchemaHelper](https://dbdiagram.io)

## Setting PostgresDB for development with docker
```bash
# make sure to run in root folder
docker-compose up -d

# getting a postgres shell inside postgres container with root user
docker exec -it postgresdb psql -U root

# run all the scripts inside BankService.sql to initialize db
```

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License
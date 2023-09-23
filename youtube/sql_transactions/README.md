# Golang Transactions with PostgreSQL

- [Video Link](https://www.youtube.com/watch?v=BLr2V6zB5k4&ab_channel=MarioCarrion)
- Run the following SQL Query in PostgreSQL Docker Container to get started
  - In `users.sql`

## Initialise Postgres with Docker

```docker
docker run -d \
-e POSTGRES_HOST_AUTH_METHOD=trust \
-e POSTGRES_USER=user \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=dbname \
-p 5433:5432 \
postgres:12.5-alpine

docker ps // to get container id

docker exec -it <container_id> psql -U user -d dbname

// You should see this:
dbname=#

// Create Table users
dbname=# CREATE TABLE users(
dbname(# name VARCHAR NOT NULL,
dbname(# birth_year SMALLINT NULL DEFAULT 0);
CREATE TABLE

dbname=# INSERT INTO users(name, birth_year) VALUES('Leon', 1997);
dbname=# SELECT * FROM users;
```

## SQL Transactions

- `Begin`, `Commit`, `Rollback`

```go
tx, err := db.BeginTx(context.Background(), nil)
tx.Commit
tx.Rollback
```



# REST web service with Golang

## DB schema migration.

### Pre-req:
#### 1. Install golang-migrate,
```sh
    brew install golang-migrate

    migrate create -ext sql -dir db/migration -seq init_schema
    /Users/dillipnayak/learning/Golang/simple_bank/db/migration/000001_init_schema.up.sql
    /Users/dillipnayak/learning/Golang/simple_bank/db/migration/000001_init_schema.down.sql

    -ext , file extension [ ex- .sql ]
    -dir , directory where migrate schema files will be placed [ex- db/migration]
    -seq , to denote a sequential schema migration version
    init_schema , name of the schema

    000001_init_schema.up.sql -  for forward changes to schema
    000001_init_schema.down.sql - for rolling back the schema changes
    
```


#### 2. create a make file to automate the db creation and schema migration tasks via make commands
File name `makefile` should be present where the command would be executed.
```md
    Content of makefile

postgres:
	docker run --name=postgres12  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --user=root --owner=root simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
dropdb:
	docker exec -it postgres12 dropdb simple_bank

.PHONY: createdb dropdb postgres migrateup migratedown

```

#### 3. Postgres container run and DB creation via make commands 
```sh
    make postgres
    make createdb

```
#### 4.  Generate schema with dbdiagram.io and run migrate 
- https://dbdiagram.io/d/6491147802bd1c4a5ebdfac1
![schema](./resources/dbdiangram.png)
- export the schema to a .sql file
- copy the content to `simple_bank/db/migration/000001_init_schema.up.sql` file
    ```sh

        migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

        2023/06/21 10:19:21 Start buffering 1/u init_schema
        2023/06/21 10:19:21 Read and execute 1/u init_schema
        2023/06/21 10:19:22 Finished 1/u init_schema (read 7.039847ms, ran 27.777225ms)
        2023/06/21 10:19:22 Finished after 41.888964ms
        2023/06/21 10:19:22 Closing source and database

    ```
- refresh the tables in tableplus to view the changes
- add `migrateup` and `migratedown` tasks to make file







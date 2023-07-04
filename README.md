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
    $make postgres
    $make createdb

```
#### 4.  Generate schema with dbdiagram.io and run migrate 
- https://dbdiagram.io/d/6491147802bd1c4a5ebdfac1
<img src="./resources/dbdiagram.png" style="float: left; margin-right: 10px;" />

- export the schema to a .sql file
- copy the content to `simple_bank/db/migration/000001_init_schema.up.sql` file
    ```sh

        $migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

        2023/06/21 10:19:21 Start buffering 1/u init_schema
        2023/06/21 10:19:21 Read and execute 1/u init_schema
        2023/06/21 10:19:22 Finished 1/u init_schema (read 7.039847ms, ran 27.777225ms)
        2023/06/21 10:19:22 Finished after 41.888964ms
        2023/06/21 10:19:22 Closing source and database

    ```
- refresh the tables in tableplus to view the changes
- add `migrateup` and `migratedown` tasks to make file
## Development
### CRUD go code generation from SQL with SQLC
#### install SQLC 
- visit https://sqlc.dev
    - sqlc for postgresql, https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html
- install sqlc ` go install github.com/kyleconroy/sqlc/cmd/sqlc@latest `
- check version with `sqlc version`
- totorials , https://github.com/kyleconroy/sqlc/blob/v1.18.0/docs/tutorials/getting-started-postgresql.md
- sqlc init to         Create an empty sqlc.yaml settings file
- update sqlc.yml file 
```yml
    version: "1"
cloud:
    organization: ""
    project: ""
    hostname: ""
project:
    id: ""
packages:
    -   name: "db"
        path: "./db/sqlc"
        queries: "./db/query"
        schema: "./db/migration"
        engine: "postgresql"
        emit_json_tag: true
        emit_prepared_queries: false
        emit_interface: false
        emit_exact_table_names: false
```
- keep queries inside db/query and run ` sqlc geenrate ` command to autom generate CRUD go code 
- update the make file to run this with make command
    ```md
        sqlc:
            sqlc generate
    ```
- create a .sql file each for  `accounts , entries, transfers` tables inside db/query directory.
- visit sqlc tutorial github page to refer the query syntax. plcae the below query to acconut.sql file and run `make sqlc` to generate golang CRUD code.
    ```sql
        -- name: CreateAcconut :one
        INSERT INTO accounts(
            owner,
            balance,
            currency
        ) VALUES (
            $1, $2,$3
        ) RETURNING *;
    ```
    ```sh
        $make sqlc
    ```
- three .go files would be generated as shown below

- do the same for other CRUS operations `Get accounts` , ` Update Accounts` & ` Delete Accounts`
    ```sql
            -- name: CreateAccount :one
            INSERT INTO accounts (
                owner,
                balance,
                currency
            ) VALUES (
                $1, $2,$3
            ) RETURNING *;

            -- name: GetAccount :one
            SELECT * FROM accounts
            WHERE id = $1 LIMIT 1;

            -- name: ListAccounts :many
            SELECT * FROM accounts
            ORDER by id
            LIMIT $1
            OFFSET $2;

            -- name: UpdateAccount :one
            UPDATE accounts
            SET balance = $1
            WHERE id = $2
            RETURNING *;

            -- name: DeleteAccount :exec
            DELETE FROM accounts
            WHERE id = $1;
    ```
<img src="./resources/make_sqlc.png" style="float: left; margin-right: 10px;" />

- initialize a go module

    ```sh
        $go mod init github.com/dillipkumar47/simplebank
        go: creating new go.mod: module github.com/dillipkumar47/simplebank
        go: to add module requirements and sums:
        go mod tidy
    ```
- run `go mod tidy` to install external dependencies









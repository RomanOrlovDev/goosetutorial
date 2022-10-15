# Tutorial of golang migrations by means of pressly/goose library

This tutorial serves to show how to use migrations in GoLang by means of "github.com/pressly/goose" library.

Run postgres in docker container:
```
docker run --name goose-postgresql -d --rm \
-p 5436:5432 \
-e POSTGRES_PASSWORD=1234 \
postgres:11
```
I choose 11th version, but it is not that important for the tutorial.

Create go file with migration:
```
goose postgres "postgresql://postgres:1234@localhost:5436/postgres?sslmode=disable" create init_user_table go  
```
Note:
- instead of postgres you can use another dialects. For this you have to check dialect.go file
- there are two options to use in the end of this command: go|sql. Files with corresponding extensions will be created.
- init_user_table - is the name of the file for creation which will be prefixed after creation
- you have to be in migrations folder before executing the command

Afterwards you have to fill created files migration manually for *up* and *down* commands. Also, you have to use tx.Exec manually and handle an error there.

Then you have to create binary with created migration files. For that you have to write go code where you specify directory to migrations files. In my case it is main.go. Then go to the directory level where this go handler added and execute:

```
go build -o my_binary_for_migrations main.go 
```

Run created binary (goose)
```
./my_binary_for_migrations postgres "postgresql://postgres:1234@localhost:5436/postgres?sslmode=disable" status
```
```
./my_binary_for_migrations postgres "postgresql://postgres:1234@localhost:5436/postgres?sslmode=disable" up
```

Finally, you can check that migrations are applied by:

`
docker exec -it goose-postgresql psql -U postgres
`

in psql console to show all created tables as well as goose tables that were created automatically for keeping data for migrations:

`
\d
`

check *goose_db_versions* to check applied migrations operations:

``
select * from goose_db_version;
``
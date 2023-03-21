# outsert

A tool for converting SQL result sets into INSERT statements.

Outsert is a useful tool for developers working with SQL databases. It simplifies the process of converting SQL result sets into INSERT statements, which can be used to insert data into another database or to recreate a database from a backup.

The example you provided shows how to use outsert in two different scenarios. The first command retrieves all the data from the "users" table in a PostgreSQL database using the psql command and pipes the output to outsert. Outsert then converts the result set into INSERT statements that can be executed on another database or used to recreate the "users" table.

The second command shows how to use outsert with a specific database URL and SQL query. This can be useful if you want to retrieve data from a remote database and insert it into a local database or vice versa.

Overall, outsert is a handy tool for developers who work with SQL databases and need a quick and easy way to convert result sets into INSERT statements.

## Example

```
$ psql -c 'select * from users' | outsert
```

```
$ DATABASE_URL=postgres://user@localhost:5432/db outsert 'select * from users'
```

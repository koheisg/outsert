# outsert

A tool for converting SQL result sets into INSERT statements.

## Example

```
$ psql -c 'select * from users' | outsert
```

```
$ DATABASE_URL=postgres://user@localhost:5432/db outsert 'select * from users'
```

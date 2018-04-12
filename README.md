# pgbytea - Test file blob store in Go

This small excercise tests out the blob storage of postgres in combination with golang. It provides a small binary that is doing the work.

 - `init`: ensures the database has the right table
 - `store`: stores a file in postgres
 - `load`: loads a file from postgres

## Usage

```
export PG_CONN='user=chartmann dbname=chartmann sslmode=disable'
pgbytea init
pgbytea add hello hello.md
pgbytea read hello > test.md
pgbytea delete hello
```


## References

- http://go-database-sql.org
- https://godoc.org/github.com/lib/pq
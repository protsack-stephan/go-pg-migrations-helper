# GO PostgreSQL migrations helper

This is and addition to awesome [https://github.com/robinjoseph08/go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) that helps you to write migrations without SQL.

Example usage:


```go
table := pgmigrations.Table{
  Name: "ninjas",
  Constraints: map[pgmigrations.Constraint][]string{
    pgmigrations.ConstraintPrimaryKey: []string{
      pgmigrations.Columns([]string{
        "user_id",
        "ninja_type",
        "ninja_id",
      }),
    },
  },
  Columns: []pgmigrations.Column{
    {
      Name: "user_id",
      Type: "int not null references users(id)",
    },
    {
      Name: "ninja_type",
      Type: "varchar(25) not null",
    },
    {
      Name: "ninja_id",
      Type: "bigint not null",
    },
    {
      Name: "updated_at",
      Type: "timestamp with time zone not null",
    },
    {
      Name: "created_at",
      Type: "timestamp with time zone not null",
    },
  },
}

// then by calling 
fmt.Println(table.Create()) // you will get SQL generated CREATE TABLE ninjas ...
// after you can just put that SQL into db.Exec and that's it
```
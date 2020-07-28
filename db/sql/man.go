package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	//db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=test sslmode=disable")
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}
	//fmt.Sprint("insert into %s (%s %s) values ($1,$2) returning id", table, column1, column2)
	//stmt := "insert into test (column_1,column_2) values ($1,$2)"
	stmt := "INSERT INTO test(column_1, column_2) VALUES(?,?)"
	prepare, err := db.Prepare(stmt)
	if err != nil {
		panic(err)
	}
	defer prepare.Close()

	rows, err := prepare.Query(123.456, 987.654)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}

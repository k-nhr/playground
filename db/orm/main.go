package main

import (
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driver   = "mysql"
	user     = "root"
	password = "root"
	dbname   = "test"
)

type Test struct {
	Column_1 float32
	Column_2 float32
}

func main() {
	stmt := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err := gorm.Open(driver, stmt)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	num := Test{Column_1: 111.456, Column_2: 222.345}
	if record := db.NewRecord(num); !record {
		panic(fmt.Errorf("failed to create record"))
	}
	db.Create(&num)
}

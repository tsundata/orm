package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tsundata/orm"
)

func main() {
	e, _ := orm.NewEngine("sqlite3", "demo.db")
	defer e.Close()
	s := e.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Tom").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}

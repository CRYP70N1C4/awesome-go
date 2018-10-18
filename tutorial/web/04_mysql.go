package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func testMySQL()  {
	db, _ := sql.Open("mysql", "root:123456@tcp(10.235.118.205:3306)/xgame")
	defer db.Close()
	rows, _ := db.Query("show tables")

	for rows.Next() {
		var tbName string
		rows.Scan(&tbName)
		fmt.Println(tbName)
	}
}

//go get -u github.com/go-sql-driver/mysql
// go install github.com/go-sql-driver/mysql
func main() {
	testMySQL()

}

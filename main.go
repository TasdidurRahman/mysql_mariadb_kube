package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db,err := sql.Open("mysql","root:rootpwd1@(localhost:3306)/")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

}
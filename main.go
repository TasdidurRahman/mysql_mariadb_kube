package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db,err := sql.Open("mysql","root:root@(127.0.0.1:3308)/")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	allDB := listDB(db)
	for _,d := range allDB{
		fmt.Println(d)
	}
}

func listDB(db *sql.DB)(array []Database) {
	res,err := db.Query("show databases;")
	if err != nil{
		//handle error
		fmt.Println(err)
	}
	for res.Next(){
		var a string
		res.Scan(&a)
		fmt.Println(a)
		if is_systemDB(a) == true{
			continue
		}
		array = append(array,getDB(a,db))
	}
	return array
}

func is_systemDB(dbname string) bool {

	switch dbname {
	case "mysql":
		return true
	case "sys":
		return true
	case "information_schema":
		return true
	case "performance_schema":
		return true
	}
	return false
}

/*
https://pkg.go.dev/database/sql
https://mariadb.com/kb/en/create-database/
https://mariadb.com/kb/en/drop-database/
https://mariadb.com/kb/en/alter-database/
*/
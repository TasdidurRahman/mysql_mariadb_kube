package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	//docker run -p 3307:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:latest
	db,err := sql.Open("mysql","root:root@(127.0.0.1:3307)/")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	db1 := Database{
		Name: "db1_name",
		CharacterSet: utf8mb3,
	}
	db1.createDB(db)

	res , err:= db.Query("show databases;")
	fmt.Println("\n all databases======= \n")
	for res.Next(){
		var s string
		res.Scan(&s)
		fmt.Println(s)
	}
	fmt.Println("\n============\n")

	res, err = db.Query("show create database db1_name;")
	var x,y string
	res.Next()
	res.Scan(&x,&y)
	fmt.Println("\n" + y + "'\n")

	db2 := Database{
		Name: "db1_name",
		CharacterSet: big5,
	}

	db1.alterDB(db,db2)

	res, err = db.Query("show create database db1_name;")
	res.Next()
	res.Scan(&x,&y)
	fmt.Println("\n" + y + "'\n")

	db1.deleteDB(db)

	res , err = db.Query("show databases;")
	fmt.Println("\n all databases======= \n")
	for res.Next(){
		var s string
		res.Scan(&s)
		fmt.Println(s)
	}
	fmt.Println("\n============\n")

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

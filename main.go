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

	//db.Exec("create database connect")

	columobj1 := Column{
		Name: "name",
		DataType: varchar10,
		AllowNull: no,
	}

	columobj2 := Column{
		Name: "age",
		DataType: integer,
		AllowNull: no,
	}

	columobj3 := Column{
		Name: "married",
		DataType: varchar10,
		AllowNull: yes,
	}

	tabobj1 := Table{
		Name: "single",
		Columns: []Column{columobj1,columobj2,columobj3},
	}

	tabobj2 := Table{
		Name: "info",
		Columns: []Column{columobj1,columobj2},
	}

	tabobj3 := Table{
		Name: "extra",
		Columns: []Column{columobj3,columobj2},
	}


	DBobj := database{
		Name: "mongo",
		Tables: []Table{tabobj1,tabobj2},
	}

	DBobj.createDB(db)
	DBobj.createTables(db)
	DBobj.addTable(db,tabobj3)



/*
	adb := database{Name: "a"}
	db.Exec(adb.deleteDB())
*/



}
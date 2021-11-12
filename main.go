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


	olddb := Database{
		Name:         "bangla",
		CharacterSet: "",
		Collation:    "",
		Encryption:   "",
		ReadOnly:     0,
	}

	//olddb.createDB(db)

	newdb := Database{
		Name:         "bangla",
		CharacterSet: big5,
		Encryption:   Y,
		ReadOnly:     enable,
	}
	olddb.alterDB(db,newdb)


}
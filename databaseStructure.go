package main

import (
	"database/sql"
	"fmt"
)

type Database struct {
	Name          string         `yaml:"name"`
	CharacterSet  charactersets  `yaml:"character_set"`
	Collation     collations     `yaml:"collation"`
	Comment 	  string         `yaml:"comment"`
}

func (d *Database)createDB(cl *sql.DB){

	//make query string ready
	query := create + database + d.Name // should it be if not exists ?
	if d.CharacterSet != nullcharset {
		query += characterset + string(d.CharacterSet)
	}
	if d.Collation != nullCollation{
		query += collation + string(d.Collation)
	}
	if d.Comment != nullComment{
		query += comment + d.Comment
	}
	query += semicolon

	// execute query
	_ , err := cl.Exec(query)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	fmt.Println(d.Name + " ==> database created...")
}

func (d *Database)deleteDB(cl *sql.DB){

	//make query string ready
	query := drop + database + d.Name + semicolon

	// execute query
	_ , err := cl.Exec(query)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	fmt.Println(d.Name + " ==> database deleted...")
}

func (d *Database)alterDB(cl *sql.DB, newdb Database){

	//dbname altering is not implemented

	// make query string ready
	query := alter + database + d.Name
	if d.CharacterSet != newdb.CharacterSet {
		d.CharacterSet = newdb.CharacterSet
		query += characterset + string(d.CharacterSet)
	}
	if d.Collation != newdb.Collation {
		d.Collation = newdb.Collation
		query += collation + string(d.Collation)
	}
	if d.Comment != newdb.Comment {
		d.Comment = newdb.Comment
		query += comment + insideSingleQuotes(d.Comment)
	}
	query += semicolon

	// execute query
	_, err := cl.Query(query)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	fmt.Println(d.Name + " ==> database altered...")
}

func (d *Database)useDB(cl *sql.DB){
	//make query string ready
	query := use + d.Name + semicolon

	//execute query
	_, err := cl.Exec(query)
	if err != nil {
		//handle error
		fmt.Println(err)
		return
	}
	fmt.Println(d.Name + " ==> using db...")
}
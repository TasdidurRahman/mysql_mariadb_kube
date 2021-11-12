package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Database struct {
	Name         string        `yaml:"name"`
	CharacterSet charactersets `yaml:"character_set"`
	Collation    collations     `yaml:"collation"`
	// mysql specific
	Encryption encryptionEnabled `yaml:"encryption"`
	ReadOnly readOnly `yaml:"read_only"`
	// mariadb specific
	Comment string `yaml:"comment"`
}

func (d *Database)createDB(cl *sql.DB){
	query := create + database + d.Name + semicolon
	cl.Exec(query)
	fmt.Println("database created...")
}

func (d *Database)deleteDB(cl *sql.DB){
	query := drop + database + d.Name + semicolon
	cl.Exec(query)
	fmt.Println("database deleted...")
}

func (d *Database)alterDB(cl *sql.DB, newdb Database){
	query := alter + database + d.Name
	if d.CharacterSet != newdb.CharacterSet {
		d.CharacterSet = newdb.CharacterSet
		query += characterset + string(d.CharacterSet)
	}
	if d.Collation != newdb.Collation {
		d.Collation = newdb.Collation
		query += collation + string(d.Collation)
	}
	// mysql
	if d.ReadOnly != newdb.ReadOnly {
		d.ReadOnly = newdb.ReadOnly
		query += readonly + strconv.Itoa(int(d.ReadOnly))
		//fmt.Println(int(d.ReadOnly))
	}
	// mysql
	if d.Encryption != newdb.Encryption {
		d.Encryption = newdb.Encryption
		query += encryption + string(d.Encryption)
	}
	//mariadb
	if d.Comment != newdb.Comment {
		d.Comment = newdb.Comment
		query += comment + insideSingleQuotes(d.Comment)
	}

	query += semicolon
	fmt.Println(query)
	cl.Exec(query)
	fmt.Println("database altered...")
}

func (d *Database)useDB(cl *sql.DB){
	query := use + d.Name + semicolon
	cl.Exec(query)
}






/*
func (d * database)createTables(cl *sql.DB){
	d.useDB(cl)
	for _,table := range d.Tables {
		cl.Exec(table.create())
	}
	fmt.Println("tables created...")
}

func (d *database)addTable(cl *sql.DB,t Table) {
	d.Tables = append(d.Tables,t)
	d.useDB(cl)
	cl.Exec(t.create())
	fmt.Println("table added...")
}

 */
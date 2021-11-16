package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type Database struct {
	Name         string        `yaml:"name"`
	CharacterSet charactersets `yaml:"character_set"`
	Collation    collations     `yaml:"collation"`
	Encryption encryptionEnabled `yaml:"encryption"`
	ReadOnly readOnly `yaml:"read_only"`

}

func (d *Database)createDB(cl *sql.DB){

	//make query string ready
	query := create + database + d.Name
	if d.CharacterSet != nullcharset {
		query += characterset + string(d.CharacterSet)
	}
	if d.Collation != nullCollation{
		query += collation + string(d.Collation)
	}
	if d.Encryption != enc_none {
		query += encryption + string(d.Encryption)
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

	//make query string ready
	query := alter + database + d.Name
	if d.CharacterSet != newdb.CharacterSet {
		d.CharacterSet = newdb.CharacterSet
		query += characterset + string(d.CharacterSet)
	}
	if d.Collation != newdb.Collation {
		d.Collation = newdb.Collation
		query += collation + string(d.Collation)
	}
	if d.Encryption != newdb.Encryption {
		d.Encryption = newdb.Encryption
		query += encryption + string(d.Encryption)
	}
	if d.ReadOnly != newdb.ReadOnly {
		d.ReadOnly = newdb.ReadOnly
		query += readonly + strconv.Itoa(int(d.ReadOnly))
	}
	query += semicolon

	//execute query
	_ , err := cl.Exec(query)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	fmt.Println("database altered...")
}

func (d *Database)useDB(cl *sql.DB){
	//make query string ready
	query := use + d.Name + semicolon

	//execute query
	_ , err := cl.Exec(query)
	if err != nil {
		//handle error
		fmt.Println(err)
		return
	}
	fmt.Println(d.Name + " ==> using db...")
}

func getDB(name string, cl *sql.DB)(ret Database){

	//make query string ready
	query := show + create + database + name + semicolon

	//execute query
	res, err := cl.Query(query)

	if err != nil {
		//handle error
		fmt.Println(err)
	}

	//process
	for res.Next(){
		var dbname, retquery string
		res.Scan(&dbname,&retquery)
		ret.Name = dbname
		fmt.Println(retquery)
		split := strings.Split(retquery," ")
		pre1 := ""
		pre2 := ""
		pre3 := ""
		for _,s := range split{
			if pre2 + space + pre1 == "CHARACTER SET"{
				ret.CharacterSet = charactersets(s)
			}
			if pre1 == "COLLATE"{
				ret.Collation = collations(s)
			}
			if pre3 + space + pre2 + space + pre1 == "READ ONLY ="{
				x,_ :=  strconv.Atoi(s)
				ret.ReadOnly = readOnly(x)
			}
			if s == "ENCRYPTION='N'"{
				ret.Encryption = N
			}
			if s == "ENCRYPTION='Y'"{
				ret.Encryption = Y
			}
			pre3 = pre2
			pre2 = pre1
			pre1 = s
		}
	}
	return ret
}

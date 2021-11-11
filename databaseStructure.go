package main

import (
	"database/sql"
	"fmt"
)

type database struct {
	Name string `yaml:"name"`
	these are cahenaslkjalsdkjfhlakjsdhf
	Tables []Table `yaml:"tables,omitempty"`
}

func (d *database)createDB(cl *sql.DB){
	query := cre + db + d.Name + end
	cl.Exec(query)
	fmt.Println("database created...")
}

func (d *database)deleteDB(cl *sql.DB){
	query := drop + db + d.Name + end
	cl.Exec(query)
	fmt.Println("database deleted...")
}

func (d *database)useDB(cl *sql.DB){
	query := use + d.Name + end
	cl.Exec(query)
}

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
package main

type columnDataType string
type allowNull string

const (
	// add space before
	integer columnDataType = " integer"
	varchar10 columnDataType = " varchar(10)"
	varchar20 columnDataType = " varchar(20)"
	no allowNull = " not null"
	yes allowNull = " null"

	// add space after
	cre string = "CREATE "
	drop string = "DROP "
	db string = "DATABASE "
	tabl string = "TABLE "
	comma string = ", "
	use string = "USE "

	// add space bothside
	on string = " ON "

	//no space
	space string = " "
	nl string = "\n"
	end string = ";"
)

func insideParenthesis(s string)(ret string){
	ret = "( " + s +" )"
	return ret
}

func insideSingleQuotes(s string)(ret string){
	ret = "'" + s +"'"
	return ret
}

func insideDoubleQuotes(s string)(ret string){
	ret = "\"" + s +"\""
	return ret
}
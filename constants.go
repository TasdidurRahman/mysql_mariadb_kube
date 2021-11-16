package main

type charactersets string
type collations string
type encryptionEnabled string
type readOnly int32

const (

	//string constants
	show 		  string = " SHOW "
	create 		  string = " CREATE "
	drop          string = " DROP "
	alter         string = " ALTER "
	database      string = " DATABASE "
	table         string = " TABLE "
	comma         string = " , "
	use           string = " USE "
	on            string = " ON "
	space         string = " "
	nl            string = "\n"
	semicolon     string = ";"
	characterset  string = " CHARACTER SET "
	collation 	  string = " COLLATE "
	readonly 	  string = " READ ONLY "
	encryption 	  string = " ENCRYPTION "
	comment 	  string = " COMMENT "
	modifyname 	  string = " MODIFY NAME "

	// charecterset constants
	nullcharset charactersets = ""
	big5     charactersets = "big5"
	dec8     charactersets = "dec8"
	cp850    charactersets = "cp850"
	hp8      charactersets = "hp8"
	koi8r    charactersets = "koi8r"
	latin1   charactersets = "latin1"
	latin2   charactersets = "latin2"
	swe7     charactersets = "swe7"
	ascii    charactersets = "ascii"
	ujis     charactersets = "ujis"
	sjis     charactersets = "sjis"
	hebrew   charactersets = "hebrew"
	tis620   charactersets = "tis620"
	euckr    charactersets = "euckr"
	koi8u    charactersets = "koi8u"
	gb2312   charactersets = "gb2312"
	greek    charactersets = "greek"
	cp1250   charactersets = "cp1250"
	gbk      charactersets = "gbk"
	latin5   charactersets = "latin5"
	armscii8 charactersets = "armscii8"
	utf8mb3  charactersets = "utf8mb3"
	ucs2     charactersets = "ucs2"
	cp866    charactersets = "cp866"
	keybcs2  charactersets = "keybcs2"
	macce    charactersets = "macce"
	macroman charactersets = "macroman"
	cp852    charactersets = "cp852"
	latin7   charactersets = "latin7"
	utf8mb4  charactersets = "utf8mb4"
	cp1251   charactersets = "cp1251"
	utf16    charactersets = "utf16"
	utf16le  charactersets = "utf16le"
	cp1256   charactersets = "cp1256"
	cp1257   charactersets = "cp1257"
	utf32    charactersets = "utf32"
	csbinary charactersets = "binary"
	geostd8  charactersets = "geostd8"
	cp932    charactersets = "cp932"
	eucjpms  charactersets = "eucjpms"


	//collation constants
	nullCollation collations = ""
	big5_chinese_ci collations = "big5_chinese_ci"
	dec8_swedish_ci collations = "dec8_swedish_ci"
	cp850_general_ci collations = "cp850_general_ci"
	hp8_english_ci collations = "hp8_english_ci"
	koi8r_general_ci collations = "koi8r_general_ci"
	latin1_swedish_ci collations = "latin1_swedish_ci"
	latin2_general_ci collations = "latin2_general_ci"
	swe7_swedish_ci collations = "swe7_swedish_ci"
	ascii_general_ci collations = "ascii_general_ci"
	ujis_japanese_ci collations = "ujis_japanese_ci"
	sjis_japanese_ci collations = "sjis_japanese_ci"
	hebrew_general_ci collations = "hebrew_general_ci"
	tis620_thai_ci collations = "tis620_thai_ci"
	euckr_korean_ci collations = "euckr_korean_ci"
	koi8u_general_ci collations = "koi8u_general_ci"
	gb2312_chinese_ci collations = "gb2312_chinese_ci"
	greek_general_ci collations = "greek_general_ci"
	cp1250_general_ci collations = "cp1250_general_ci"
	gbk_chinese_ci collations = "gbk_chinese_ci"
	latin5_turkish_ci collations = "latin5_turkish_ci"
	armscii8_general_ci collations = "armscii8_general_ci"
	utf8mb3_general_ci collations = "utf8mb3_general_ci"
	ucs2_general_ci collations = "ucs2_general_ci"
	cp866_general_ci collations = "cp866_general_ci"
	keybcs2_general_ci collations = "keybcs2_general_ci"
	macce_general_ci collations = "macce_general_ci"
	macroman_general_ci collations = "macroman_general_ci"
	cp852_general_ci collations = "cp852_general_ci"
	latin7_general_ci collations = "latin7_general_ci"
	utf8mb4_general_ci collations = "utf8mb4_general_ci"
	cp1251_general_ci collations = "cp1251_general_ci"
	utf16_general_ci collations = "utf16_general_ci"
	utf16le_general_ci collations = "utf16le_general_ci"
	cp1256_general_ci collations = "cp1256_general_ci"
	cp1257_general_ci collations = "cp1257_general_ci"
	utf32_general_ci collations = "utf32_general_ci"
	clbinary collations = "binary"
	geostd8_general_ci collations = "geostd8_general_ci"
	cp932_japanese_ci collations = "cp932_japanese_ci"
	eucjpms_japanese_ci collations = "eucjpms_japanese_ci"

	//encryptionEnabled
	Y encryptionEnabled = "'Y'"
	N encryptionEnabled = "'N'"

	//readOnly
	enable readOnly = 1
	disable readOnly = 2

	nullComment string = ""
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
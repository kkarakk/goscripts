package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	condb, errdb := sql.Open("mssql", "Server=.\\SQLExpress;Database=gts;user id=sa;password=123456")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	var (
		sqlversion string
	)
	rows, err := condb.Query("select @@version")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(sqlversion)
	}
	// rows2, err2 := condb.Query("select * from msdevice")
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	defer condb.Close()
}

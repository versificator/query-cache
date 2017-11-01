package main

import (
	"fmt"
	"database/sql"
	"github.com/kshvakov/clickhouse"
	"querycache/log"

)

func getDataFromDB(jsonRaw string) *sql.Rows {
	connect, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		log.Fatalf("%s",err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return nil
	}
    j := parseRoot(&jsonRaw)
	j.rows, err = connect.Query(j.query + " FORMAT tsv")
	if err != nil {
		log.Fatalf("Error running the query: ",err)
	}
	defer j.rows.Close()

	writeToFile(j.rows,"2017-06-13")
	// Close rows and database connection
	j.rows.Close()
	connect.Close()

	return j.rows

}
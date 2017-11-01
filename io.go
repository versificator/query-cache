package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"querycache/log"
	"database/sql"
	"path/filepath"
	"path"
)

func createFolder(folderName string){
	newPath := filepath.Join(".", folderName)
	if err := os.MkdirAll(newPath, os.ModePerm); err != nil { log.Fatalf("Cannot create hidden directory.") }
}

func writeToFile(rows *sql.Rows, fileName string){
	createFolder("cache1")
	file, err := os.Create(path.Join("cache1/", fileName))
	defer file.Close()

	out := csv.NewWriter(file)

	// Get the column names
	// columnNames := []string
	columnNames, err := rows.Columns()

	if err != nil {
		log.Fatalf("Error getting columns:", err)
	}

	// Print out the column names
	out.Write(columnNames)
	out.Flush()

	// Setup storge for SQL results as string for CSV writter
	result := make([]string, len(columnNames))

	// Setup storage for SQL Scan() results
	values := make([]interface{}, len(columnNames))
	valuePtrs := make([]interface{}, len(columnNames))

	// Iterate over the results and write CSV lines to STDOUT
	for rows.Next() {

		// Create a value pointer per column in the row
		for i, _ := range columnNames {
			valuePtrs[i] = &values[i]
		}

		// Get the column values from the row
		err = rows.Scan(valuePtrs...)

		// Abort on error
		if err != nil {
			log.Fatalf("Error getting result:", err)
		}

		// Iterate through the retrieved values and save them as strings in the
		// results array
		for i, _ := range columnNames {
			var v interface{}

			val := values[i]

			b, ok := val.([]byte)

			if (ok) {
				v = string(b)
			} else {
				v = val
			}

			if v == nil {
				v = ""
			}

			result[i] = fmt.Sprintf("%v", v)
		}

		out.Write(result)
		out.Flush()
	}

	// Check for error on rows
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error on rows", err)
	}

}
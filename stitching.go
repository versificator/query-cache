package main

import (
	"fmt"
	"encoding/json"
)


type meta struct {
	Name string    `json:"name"`
	Type string `json:"type"`
}

type columnType struct {
	Meta []meta  `json:"meta"`
}
func processing(resultJSON []byte ) {
	var columns columnType

	if err := json.Unmarshal(resultJSON, &columns); err != nil {
		fmt.Println(err)
	}

	for k, v := range columns.Meta {
		_ = k
		fmt.Println(v.Type, " ", v.Name)
	}
	fmt.Print(columns.Meta)
}

//
//func stitch(resultJSON [][]byte){
//	var result []byte
//
//	for k,v :=range resultJSON{
//
//	}
//
//}



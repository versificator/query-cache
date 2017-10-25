package main

import (

"fmt"

"encoding/json"
"github.com/Jeffail/gabs"
"time"
"reflect"
"bytes"
	"strings"
)

type Query struct {
	HourGroup     bool      `json:"hourGroup"` //flag, is hourly breakdown required
	DateTimeStart time.Time `json:"dateTimeStart"`
	DateTimeEnd   time.Time `json:"dateTimeEnd"`
	Select        []string  `json:"select"`
	From          []string  `json:"from"`
	PreWhere      []string  `json:"preWhere"`
	Where         []string  `json:"where"`
	Range         []string  `json:"range"`
	Dimension     []string  `json:"groupBy"` //groupBy columns
}

var JSONAttr = []string {"$select", "$from", "$preWhere", "$where", "$groupBy", "$having", "$globalRangePreWhere", "$globalRangeWhere"}


func jsonToQuery(jsonQuery string) string {
	var query interface{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	byteValue := []byte(jsonQuery)

	json.Unmarshal(byteValue, &query)

	m := query.(map[string]interface{})
	var result string

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	return result
}

func  parseRoot(json *string) {
	sample := []byte(*json)
	jsonParsed, _ := gabs.ParseJSON(sample)
	child := jsonParsed.Search("$query").Data()
	fmt.Println("\n\n\n\n\n\n", parseRootRawJSON(child))
}



func parseRootRawJSON(object interface{}) string {
	result := make(map[string]string,8)

	if reflect.TypeOf(object).String() == "string" {
		fmt.Println(object.(string))
	} else if reflect.TypeOf(object).String() == "[]interface {}" {
		for _, value := range object.([]interface{}) {
			fmt.Println(value.(string))
		}
	} else if reflect.TypeOf(object).String() == "map[string]interface {}" {
		objectMap := object.(map[string]interface{})

		for keyLevel1, valueLevel1 := range objectMap {
			fmt.Println("\n", keyLevel1, " LEVEL1  ", valueLevel1)

			if keyLevel1 == "$where" || keyLevel1 == "$preWhere" || keyLevel1 == "$globalRangePreWhere" || keyLevel1 == "$globalRangeWhere" {
				for keyLevel2, valueLevel2 := range valueLevel1.(map[string]interface{}) {
					fmt.Println(keyLevel2, " $$LEVEL2 - WHERE$$  ", valueLevel2)

					result[keyLevel1] = concat( " 1=1 ")
					if keyLevel2 == "$and" {
						for keyLevel3, valueLevel3 := range valueLevel2.(map[string]interface{}) {
							fmt.Println(keyLevel3, " $$LEVEL3 - WHERE$$  ", valueLevel3)
							if keyLevel3 == "$range" {
								if keyLevel1 == "$where" {result[keyLevel1] = concat(result[keyLevel1], " and $globalRangeWhere ")
								} else if keyLevel1 == "$preWhere" {result[keyLevel1] = concat(result[keyLevel1], " and $globalRangePreWhere ")}
									fmt.Println("==============================", keyLevel1,"  ",result[keyLevel1])
							} else if keyLevel3 == "$in" {
								if reflect.TypeOf(valueLevel3).String() == "map[string]interface {}" {
									for keyLevel4, valueLevel4 := range valueLevel3.(map[string]interface{}) {
										fmt.Println(keyLevel4, " $$LEVEL5 - WHERE$$  ", valueLevel4)
										if  reflect.TypeOf(valueLevel4).String() == "map[string]interface {}"{
											for keyLevel5, valueLevel5 := range valueLevel4.(map[string]interface{}){
												if keyLevel5 =="$query" {result[keyLevel1] = concat(result[keyLevel1]," and ", keyLevel4  ," in (",  parseRootRawJSON(valueLevel5),")")}
										}
										} else if reflect.TypeOf(valueLevel4).String() == "string"{
											result[keyLevel1] = concat(result[keyLevel1]," and ", keyLevel4," in (", valueLevel4.(string),")")
										}
									}
									fmt.Println("==============================", keyLevel1, "  ", result[keyLevel1])
								}
							} else if reflect.TypeOf(valueLevel3).String() == "map[string]interface {}" {
								for keyLevel4, valueLevel4 := range valueLevel3.(map[string]interface{}) {
									fmt.Println(keyLevel4, " $$LEVEL4 - WHERE$$  ", valueLevel4)
                                      if reflect.TypeOf(valueLevel4).String() == "map[string]interface {}" {
										for keyLevel5, valueLevel5 := range valueLevel4.(map[string]interface{}) {
											result[keyLevel1] = concat(result[keyLevel1], " and ", keyLevel5, keyLevel4, valueLevel5.(string))
											fmt.Println("==============================", keyLevel1, "  ", result[keyLevel1])
										}
									} else {
										result[keyLevel1] = concat(result[keyLevel1], " and ", keyLevel4, keyLevel3, valueLevel4.(string))
										fmt.Println("==============================", keyLevel1, "  ", result[keyLevel1])
									}

								}
							}
						}
					}

				}

			} else if keyLevel1 == "$from" {
				fmt.Println(keyLevel1, " LEVEL2 - FROM ", valueLevel1)
				for _, y := range valueLevel1.([]interface{}) {
					if reflect.TypeOf(y).String() == "string" {
						result[keyLevel1] = concat(result[keyLevel1],y.(string))
					}
				}
				fmt.Println("==============================", keyLevel1, result[keyLevel1])

			} else if keyLevel1 == "$groupBy" {
				result[keyLevel1] = concat(result[keyLevel1]," ")
				for x, y := range valueLevel1.([]interface{}) {
					if x != 0 {
						result[keyLevel1] = concat(result[keyLevel1],",")
					}
					result[keyLevel1] = concat(result[keyLevel1], y.(string))
				}
				fmt.Println("==============================", keyLevel1, result[keyLevel1])
			} else if keyLevel1 == "$select" {
				for x, y := range valueLevel1.([]interface{}) {
					if reflect.TypeOf(y).String() == "string" {
						if x != 0 {
							result[keyLevel1] = concat(result[keyLevel1], ",")
						}
						result[keyLevel1] = concat(result[keyLevel1], y.(string))
					}
				}
				fmt.Println("==============================", keyLevel1, result[keyLevel1])
			} else if keyLevel1 == "$having" {
				for keyLevel2, valueLevel2 := range valueLevel1.(map[string]interface{}) {
					fmt.Println(keyLevel2, " LEVEL2 - HAVING ", valueLevel2)
					if keyLevel2 == "$or" {
						for keyLevel3, valueLevel3 := range valueLevel2.(map[string]interface{}) {
							fmt.Println(keyLevel3, " LEVEL3 - HAVING ", valueLevel3)
							for keyLevel4, valueLevel4 := range valueLevel3.(map[string]interface{}) {
								fmt.Println(keyLevel4, " LEVEL4 - HAVING ", valueLevel4)
								result[keyLevel1] = strings.TrimLeft(concat(result[keyLevel1], " or ", keyLevel3, " ", keyLevel4, " ", valueLevel4.(string))," or ")
							}
						}
						fmt.Println("==============================", keyLevel1, result[keyLevel1])
					}
				}
			}

		}
	}
	return toString(result)
}

func concat(values ...string) string {
	var buffer bytes.Buffer
	for _, s := range values {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func toString(buf map[string]string) string {
	var result string
//	fmt.Println(buf)
	for _,y := range JSONAttr{
		switch y {
		case "$select":
			result = concat(result, "select ", buf[y])
		case "$from":
			result = concat(result, " \nfrom ", buf[y])
		case "$preWhere":
			if buf[y] !="" {result = concat(result, " \nprewhere ", buf[y])}
		case "$where":
			if buf[y] !="" {result = concat(result, " \nwhere ", buf[y])}
		case "$groupBy":
			if buf[y] !="" {result = concat(result, " \ngroup by ", buf[y])}
		case "$having" :
			if buf[y] !="" {result = concat(result, " \nhaving ", buf[y])}
		case "$globalRangePreWhere":
			{
				fmt.Println("\n\n\n\n\n\nRANGE!!!!!!!!!!!!!", buf[y])
				result = strings.Replace(result, y, buf[y], len(result))
			}
		case "$globalRangeWhere":
			{
				fmt.Println("\n\n\n\n\n\nRANGE!!!!!!!!!!!!!", buf[y])
				result = strings.Replace(result, y, buf[y], len(result))
			}
		}
		fmt.Println(result)
	}
	return result
}






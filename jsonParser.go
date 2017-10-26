package main

import (
"fmt"
"github.com/Jeffail/gabs"
"reflect"
"bytes"
"strings"
"log"
)


var JSONAttr = []string {"$select", "$from", "$preWhere", "$where", "$groupBy", "$having", "$globalRangePreWhere", "$globalRangeWhere"}

func  parseRoot(json *string) {
	sample := []byte(*json)
	jsonParsed, _ := gabs.ParseJSON(sample)
	child := jsonParsed.Search("$query").Data()
	fmt.Println("\n\n\n\n\n\n", toString(parseRootRawJSON(child)))
}


func parseRootRawJSON(object interface{}) map[string]string{
	result := make(map[string]string,8)

	if reflect.TypeOf(object).String() == "string" {
		log.Fatal(object.(string))
	} else if reflect.TypeOf(object).String() == "[]interface {}" {
		for _, value := range object.([]interface{}) {
			log.Fatal(value.(string))
		}
	} else if reflect.TypeOf(object).String() == "map[string]interface {}" {
		objectMap := object.(map[string]interface{})
		for keyLevel1, valueLevel1 := range objectMap {
			log.Println(keyLevel1, " LEVEL1  ", valueLevel1)
			if keyLevel1 == "$where" || keyLevel1 == "$preWhere" || keyLevel1 == "$globalRangePreWhere" || keyLevel1 == "$globalRangeWhere" {
				for keyLevel2, valueLevel2 := range valueLevel1.(map[string]interface{}) {
					log.Println(keyLevel2, " LEVEL2(WHERE) ", valueLevel2)
					if keyLevel2 == "$and" {
						for keyLevel3, valueLevel3 := range valueLevel2.(map[string]interface{}) {
							log.Println(keyLevel3, " LEVEL2(WHERE) ", valueLevel3)
							if keyLevel3 == "$range" {
								if keyLevel1 == "$where" {result[keyLevel1] = concat(result[keyLevel1], " $globalRangeWhere and ")
								} else if keyLevel1 == "$preWhere" {result[keyLevel1] = concat(result[keyLevel1], " $globalRangePreWhere and ")}
								log.Println("Result:                          ", keyLevel1,"  ",result[keyLevel1])
							} else if keyLevel3 == "$in" {
								if reflect.TypeOf(valueLevel3).String() == "map[string]interface {}" {
									for keyLevel4, valueLevel4 := range valueLevel3.(map[string]interface{}) {
										log.Println(keyLevel4, " LEVEL5(WHERE) ", valueLevel4)
										if  reflect.TypeOf(valueLevel4).String() == "map[string]interface {}"{
											for keyLevel5, valueLevel5 := range valueLevel4.(map[string]interface{}) {
												if keyLevel5 == "$query" {
													result[keyLevel1] = concat(result[keyLevel1], keyLevel4, " in (", toString(parseRootRawJSON(valueLevel5.(interface{}))), ") and ")
												}
												log.Println("Nested query :", result[keyLevel1])
											}
										} else if reflect.TypeOf(valueLevel4).String() == "string"{
											result[keyLevel1] = concat(result[keyLevel1], keyLevel4," in (", valueLevel4.(string),") and ")
										}
									}
									log.Println("Result:                          ", keyLevel1, "  ", result[keyLevel1])
								}
							} else if reflect.TypeOf(valueLevel3).String() == "map[string]interface {}" {
								for keyLevel4, valueLevel4 := range valueLevel3.(map[string]interface{}) {
									log.Println(keyLevel4, " LEVEL4(WHERE) ", valueLevel4)
                                      if reflect.TypeOf(valueLevel4).String() == "map[string]interface {}" {
										for keyLevel5, valueLevel5 := range valueLevel4.(map[string]interface{}) {
											result[keyLevel1] = concat(result[keyLevel1], keyLevel5, keyLevel4, valueLevel5.(string)," and ")
											log.Println("Result:                          ", keyLevel1, "  ", result[keyLevel1])
										}
									} else {
										result[keyLevel1] = concat(result[keyLevel1], keyLevel4, keyLevel3, valueLevel4.(string), " and ")
										  log.Println("Result:                          ", keyLevel1, "  ", result[keyLevel1])
									}
								}
							}
						}
						log.Println("TRIMRIGHT")
						result[keyLevel1] = strings.TrimRight(result[keyLevel1]," and ")
					}

				}

			} else if keyLevel1 == "$from" {
				log.Println(keyLevel1, " LEVEL2(FROM) ", valueLevel1)
				for _, y := range valueLevel1.([]interface{}) {
					if reflect.TypeOf(y).String() == "string" {
						result[keyLevel1] = concat(result[keyLevel1],y.(string))
					}
				}
				log.Println("Result:                          ", keyLevel1, result[keyLevel1])

			} else if keyLevel1 == "$groupBy" {
				result[keyLevel1] = concat(result[keyLevel1]," ")
				for x, y := range valueLevel1.([]interface{}) {
					if x != 0 {
						result[keyLevel1] = concat(result[keyLevel1],",")
					}
					result[keyLevel1] = concat(result[keyLevel1], y.(string))
				}
				log.Println("Result:                          ", keyLevel1, result[keyLevel1])
			} else if keyLevel1 == "$select" {
				for x, y := range valueLevel1.([]interface{}) {
					if reflect.TypeOf(y).String() == "string" {
						if x != 0 {
							result[keyLevel1] = concat(result[keyLevel1], ",")
						}
						result[keyLevel1] = concat(result[keyLevel1], y.(string))
					}
				}
				log.Println("Result:                          ", keyLevel1, result[keyLevel1])
			} else if keyLevel1 == "$having" {
				for keyLevel2, valueLevel2 := range valueLevel1.(map[string]interface{}) {
					log.Println(keyLevel2, " LEVEL2(HAVING) ", valueLevel2)
					if keyLevel2 == "$or" {
						for keyLevel3, valueLevel3 := range valueLevel2.(map[string]interface{}) {
							log.Println(keyLevel3, " LEVEL3(HAVING) ", valueLevel3)
							for keyLevel4, valueLevel4 := range valueLevel3.(map[string]interface{}) {
								log.Println(keyLevel4, " LEVEL4(HAVING) ", valueLevel4)
								result[keyLevel1] = strings.TrimLeft(concat(result[keyLevel1], " or ", keyLevel3, " ", keyLevel4, " ", valueLevel4.(string))," or ")
							}
						}
						log.Println("Result:                          ", keyLevel1, result[keyLevel1])
					}
				}
			}

		}
	}
	return result
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
				if buf[y] != "" {
					result = strings.Replace(result, y, buf[y], len(result))
				}
			}
		case "$globalRangeWhere":
			{
				if buf[y] != "" {
					log.Println("Replace ", y, " ", result, " ", buf[y])
					result = strings.Replace(result, y, buf[y], len(result))
				}
			}
		}
		log.Println(result)
	}
	return result
}






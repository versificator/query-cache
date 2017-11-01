package main

import (
"github.com/Jeffail/gabs"
"reflect"
"bytes"
"strings"
"log"
"encoding/json"
"database/sql"

	"fmt"
)

type dateRange struct{
	startDate string
	endDate string
	startTimestamp uint8
	endTimestamp uint8
}

type JSON struct {
	rawJSON    *string
	query      string
	dateRange  dateRange
	rows *sql.Rows
	fileName   string
	queryHash  cache
}

var JSONAttr = []string {"$select", "$from", "$preWhere", "$where", "$groupBy", "$having", "$globalRangePreWhere", "$globalRangeWhere"}

func getDateRange(g map[string]string) dateRange {
	fmt.Println("asdf            ",g["$globalRangeWhere"])

	l := dateRange{g["globalRangeWhere"],"",0,0}
	return l
}

func  parseRoot(json *string) JSON {
	j := JSON{json,"",dateRange{"","",0,0},nil,"test",cache{}}
	sample := []byte(*json)
	jsonParsed, _ := gabs.ParseJSON(sample)
	child := jsonParsed.Search("$query").Data()
	parsedJSON := parseRootRawJSON(child)
	j.query = toString(parsedJSON)
	j.dateRange = getDateRange(parsedJSON)
	return j
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
							log.Println(keyLevel3, " LEVEL3(WHERE) ", valueLevel3)
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

//type object struct{}
//
//type Querier interface {
//	ToQuery() string
//}
//
//func (o object) ToQuery() string {
//	return ""
//	}
////map["columns"] slice Querier
//
//type SelectExprs []SelectExpr
//
//type SelectExpr interface {
//	iSelectExpr()
//	SQLNode
//}
//
//type AST struct {
//
//	SelectExprs SelectExprs
//	From        TableExprs
//	Where       *Where
//	GroupBy     GroupBy
//	Having      *Where
//	OrderBy     OrderBy
//	Limit       *Limit
//	Lock        string
//}

//type AST struct {
//	columns  []Querier
//	from     []Querier
//	where    []Querier
//	preWhere []Querier
//}

//
//func (a *AST) compile() string {
//	var result string
//	for _, c := range a.columns {
//		result += c.ToQuery()
//	}
//	return result
//}
//
//type condition struct {
//	field    string
//	operator string
//	value    string
//}
//
//func prepend(a []condition, b *condition) []condition{
//	var result []condition
//	result = append(result, condition{"","",""})
//	return result
//}
//func (a *AST) toString() string {
//	return " "
//}
//func parseRootRawJSON_2() {
//	sql = `selectclo1, col2from(select * from table2)wherecondition`
//
//	c1 := &condition{field: "col1", operator: "=", value: "12"}
//	a := &AST{}
//	startDateCondition := &condition{field: "EventHour", operator: "=", value: "12"}
//	endDateCondition := &condition{field: "EventDate", operator: "=", value: "today()"}
//
//	a.where = prepend(a.where, startDateCondition)
//	a.where = prepend(a.where, endDateCondition)
//
//	sql := s.String()
//
//	string{}
//	reg := map[string][]string
//}
//	func parse() {
//		var root string
//	tree := type Tree struct{}
//
//	for {
//		token, ok := tree.getToken()
//
//		select col1
//			col2
//			from
//			table
//			where
//
//			if !ok {
//				return
//			}
//			switch token {
//			case "$from":
//				root = "$from"
//			case "$select":
//				root = "$select"
//				// do
//			case "$where", "$preWhere", "$globalRangePreWhere", "$globalRangeWhere":
//				// do
//			default:
//				if root != nil {
//					reg[root] = []string{token}
//					// add to current root level} else {
//					// add to top level}}}}
//
//				}
//			}
//		}

//remove range predicates from JSON
func cleanJSON(jsonRaw *string) []byte {
	sample := []byte(*jsonRaw)
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(sample), &jsonMap)
	if err != nil {
		panic(err)
	}

	jsonMap["$query"].(map[string]interface{})["$globalRangePreWhere"] = nil
	jsonMap["$query"].(map[string]interface{})["$globalRangeWhere"] = nil
	result, err := json.Marshal(jsonMap)
	if err != nil {
		panic(err)
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
		fmt.Print(result)
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






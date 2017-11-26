package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type rangeHour struct{
	StartHour string `json:"$startHour"`
	EndHour string `json:"$endHour"`
}

type rangeDay struct{
	StartDay string `json:"$startDay"`
	EndDay string `json:"$endDay"`
}

type query struct {
	Query      string    `json:"$query"`
	RangeDay   rangeDay `json:"RangeDay"`
	RangeHour  rangeHour `json:"RangeHour"`
}

func  parse(rawJSON string) query {
	var result query
	bytes := []byte(rawJSON)

	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println(err)
	}
	return result
}



func toString(buf query) string {
	r := strings.NewReplacer("$startHour", buf.RangeHour.StartHour,
		"$endHour", buf.RangeHour.EndHour,
		"$startDay", buf.RangeDay.StartDay,
		"$endDay", buf.RangeDay.EndDay)

	result := r.Replace(buf.Query)
	return result

	}







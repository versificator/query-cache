package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type query struct {
	Query      string    `json:"$query"`
	StartHour uint32 `json:"$startHour,string"`
	EndHour uint32 `json:"$endHour,string"`
	StartDay time.Time `json:"$startDay,string"`
	EndDay time.Time `json:"$endDay,string"`
	GroupBy []string `json:"$groupby"`
}

func  parse(rawJSON []byte) query {
	var result query
	if err:=json.Unmarshal(rawJSON,&result); err !=nil {
		fmt.Println(err)
	}
	return result
}

func toString(buf query) string {
	r := strings.NewReplacer("$startHour", fmt.Sprint(buf.StartHour),
		"$endHour", fmt.Sprint(buf.EndHour),
		"$startDay",  buf.StartDay.Format("2006-01-02"),
		"$endDay", buf.EndDay.Format("2006-01-02"))

	result := r.Replace(buf.Query)
	return result
}









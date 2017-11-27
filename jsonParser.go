package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"strconv"

)

type RangeDate struct{
	StartHour uint32 `json:"$startHour,string"`
	EndHour uint32 `json:"$endHour,string"`
	StartDay time.Time `json:"$startDay,string"`
	EndDay time.Time `json:"$endDay,string"`
}


type query struct {
	Query      string    `json:"$query"`
	Filter   RangeDate `json:"RangeDate"`
}

func  parse(rawJSON string) query {
	var result query
	bytes := []byte(rawJSON)

	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println(err)
	}
	return result
}


func (l *RangeDate) UnmarshalJSON(j []byte) error {
	var rawStrings map[string]string

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "$startHour" {

			rb, err := strconv.Atoi(v)
			l.StartHour = uint32((rb))
			if err != nil {
				return err
			}
		}

		if strings.ToLower(k) == "$endHour" {
			r, err := strconv.Atoi(v)
			l.EndHour = uint32(r)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "$startDay" || strings.ToLower(k) == "$endDay" {
			t, err := time.Parse("", v)
			if err != nil {
				return err
			}
			l.StartDay = t
		}
	}

	return nil
}



func toString(buf query) string {
	r := strings.NewReplacer("$startHour", string(buf.Filter.StartHour),
		"$endHour", string(buf.Filter.EndHour),
		"$startDay",  buf.Filter.StartDay.Format("2006-01-02"),
		"$endDay", buf.Filter.EndDay.Format("2006-01-02"))

	result := r.Replace(buf.Query)
	return result
}









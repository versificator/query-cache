package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"strconv"

)

type query struct {
	Query      string    `json:"$query"`
	StartHour uint32 `json:"$startHour,string"`
	EndHour uint32 `json:"$endHour,string"`
	StartDay time.Time `json:"$startDay,string"`
	EndDay time.Time `json:"$endDay,string"`
}

func  parse(rawJSON string) query {
	var result query
	bytes := []byte(rawJSON)

	if err := result.UnmarshalJSON(bytes); err != nil {
		fmt.Println(err)
	}
	return result
}


func (l *query) UnmarshalJSON(j []byte) error {
	var rawStrings map[string]string

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "$starthour" && v != "" {
			rb, err := strconv.Atoi(v)
			l.StartHour = uint32((rb))
			if err != nil {
				return err
			}
		} else if strings.ToLower(k) == "$query" && v != "" {
			l.Query = v
		} else if strings.ToLower(k) == "$endhour" && v != "" {
			r, err := strconv.Atoi(v)
			l.EndHour = uint32(r)
			if err != nil {
				return err
			}
		} else if strings.ToLower(k) == "$startday" && v != "" {
			t, err := time.Parse("2006-01-02", v)
			if err != nil {
				return err
			}
			l.StartDay = t
		} else if strings.ToLower(k) == "$endday" && v != "" {
			t, err := time.Parse("2006-01-02", v)
			if err != nil {
				return err
			}
			l.EndDay = t
		}
	}
	return nil
}



func toString(buf query) string {
	r := strings.NewReplacer("$startHour", fmt.Sprint(buf.StartHour),
		"$endHour", fmt.Sprint(buf.EndHour),
		"$startDay",  buf.StartDay.Format("2006-01-02"),
		"$endDay", buf.EndDay.Format("2006-01-02"))

	result := r.Replace(buf.Query)
	return result
}









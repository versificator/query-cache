package main

import (
	"fmt"
	"time"
)

  //return false if range dates are not whole hours
func checkRangeDate(d query) (bool,bool){
	return d.StartHour % 3600 ==0, d.EndHour % 3600 == 0
}

func getHourRanges(d query) []query {
	var result []query

	if d.StartHour != 0 && d.EndHour != 0 {
		t1, t2 := checkRangeDate(d)
		if t1 == false {
			fmt.Errorf("start timestamp must be whole hour")
		}

		currTimestamp, nextTimestamp := d.StartHour, d.StartHour+3600
		for currTimestamp <= d.EndHour-3600 {
			result = append(result, query{d.Query, currTimestamp, nextTimestamp, timestamp2day(currTimestamp), timestamp2day(currTimestamp),[]string{""}})
			currTimestamp = nextTimestamp
			nextTimestamp = nextTimestamp + 3600
		}
		if t2 == false {
			result = append(result, query{d.Query, currTimestamp, d.EndHour, timestamp2day(currTimestamp), timestamp2day(currTimestamp),[]string{""}})
		}
	} else if d.StartHour == 0 && d.EndHour == 0 && !d.StartDay.IsZero() && !d.EndDay.IsZero() {

		a := d.StartDay;
		for {

			result = append(result, query{d.Query, 0, 0, a, a.AddDate(0, 0, 1),[]string{""}})
			if a.Equal(d.EndDay) {
				break
			}
			a = a.AddDate(0, 0, 1)
		}
	}
	return result
}

func timestamp2day(i uint32) time.Time {
	tm := time.Unix(int64(i), 0)
	rounded := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0,  time.UTC)
	return rounded
}
package main

import "fmt"

type rangeDate struct {
	startTimestamp uint64
	endTimestamp uint64
}
  //return false if range dates are not whole hours
func checkRangeDate(d rangeDate) (bool,bool){
	return d.startTimestamp % 3600 ==0, d.endTimestamp % 3600 == 0
}

func getHourRanges(d rangeDate) []rangeDate {
	var result []rangeDate
	t1, t2 := checkRangeDate(d)
	if t1 == false {
		fmt.Errorf("start timestamp must be whole hour")
	}
	if t1 == true {
		currTimestamp, nextTimestamp := d.startTimestamp, d.startTimestamp + 3600
		for currTimestamp <= d.endTimestamp - 3600 {
			result = append(result, rangeDate{currTimestamp, nextTimestamp})
			currTimestamp = nextTimestamp
			nextTimestamp = nextTimestamp + 3600
		}
		if t2 == false {
			result = append(result, rangeDate{currTimestamp, d.endTimestamp})
		}
	}
	return result
}

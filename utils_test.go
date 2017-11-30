package main

import (
	"testing"
	"time"
	"fmt"
	"sort"
)
func TestGetHourRanges_1(t *testing.T) {
	tm,err := time.Parse("02-01-2006", "23-11-2017")

	ti := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0,  time.UTC)
	if err != nil {fmt.Println(err)}

	out := []query{}
	start:= uint32(1511398800)
	for {
		out = append(out,query{StartHour:start,EndHour:start + uint32(time.Hour/1000000000),StartDay:ti,EndDay:ti})
		start=start + uint32(time.Hour/1000000000)
		if start > 1511434800 {break}
	}

	in := query{StartHour:1511398800, EndHour: 1511438400,StartDay:ti,EndDay:ti}
	if !Equals(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}
}

func TestGetHourRanges_2(t *testing.T) {
	tm,err := time.Parse("02-01-2006", "23-11-2017")

	ti := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0,  time.UTC)
	if err != nil {fmt.Println(err)}

	out := []query{}
	start:= uint32(1511398800)
	for {
		out = append(out,query{StartHour:start,EndHour:start + uint32(time.Hour/1000000000),StartDay:ti,EndDay:ti})
		start=start + uint32(time.Hour/1000000000)
		if start > 1511431200 {break}
	}
	out = append(out,query{StartHour:1511434800,EndHour:1511438000,StartDay:ti,EndDay:ti})

	in := query{StartHour:1511398800, EndHour: 1511438000,StartDay:ti,EndDay:ti}
	if !Equals(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}
}


func TestGetDaysRanges_1(t *testing.T) {
	start, err := time.Parse("02-01-2006", "23-11-2017")
	if err != nil {fmt.Println(err)}
	end, err := time.Parse("02-01-2006", "30-11-2017")
	if err != nil {fmt.Println(err)}
	out := []query{}
	inc := start
	for {
		out = append(out, query{StartDay: inc, EndDay: inc.AddDate(0,0,1)})
		inc = inc.AddDate(0, 0, 1)
		if inc.After(end) {
			break
		}
	}
	in := query{StartDay: start, EndDay: end}
	if !Equals(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}
}

func Equals(a,b []query) bool {
	for i := range a {
		length := len(a) == len(b)
		queries := a[i].Query  == b[i].Query
		endDay := a[i].EndDay == b[i].EndDay
		endHour := a[i].EndHour == b[i].EndHour
		startDay := a[i].StartDay == b[i].StartDay
		startHour := a[i].StartHour == b[i].StartHour
		groupBy := len(a[i].GroupBy) == len(b[i].GroupBy)
		if groupBy {// copy slices so sorting won't affect original structs
			aGroupBy := make([]string, len(a[i].GroupBy))
			bGroupby := make([]string, len(b[i].GroupBy))
			copy(a[i].GroupBy, aGroupBy)
			copy(b[i].GroupBy, bGroupby)
			sort.Strings(aGroupBy)
			sort.Strings(bGroupby)
			for index, item := range aGroupBy {
				if item != bGroupby[index] {
					groupBy = false
				}
			}
		}
		if !(length && queries && endDay && endHour && startDay && startHour) {return false}
	}
	return true
}
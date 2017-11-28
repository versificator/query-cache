package main

import (
	"testing"
	"time"
	"fmt"
)
func TestGetHourRanges_1(t *testing.T) {
	//t.Fatalf("unexpected in out: ")
	tm,err := time.Parse("02-01-2006", "23-11-2017")

	ti := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0,  time.UTC)
	if err != nil {fmt.Println(err)}

	out := []query{query{"",1511398800, 1511402400,ti,ti},
		query{"",1511402400, 1511406000,ti,ti},
		query{"",1511406000, 1511409600,ti,ti},
		query{"",1511409600, 1511413200,ti,ti},
		query{"",1511413200, 1511416800,ti,ti},
		query{"",1511416800, 1511420400,ti,ti},
		query{"",1511420400, 1511424000,ti,ti},
		query{"",1511424000, 1511427600,ti,ti},
		query{"",1511427600, 1511431200,ti,ti},
		query{"",1511431200, 1511434800,ti,ti},
		query{"",1511434800, 1511438400,ti,ti},
	}

	in := query{"",1511398800, 1511438400,ti,ti}
	if !testEq(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}

}

func TestGetHourRanges_2(t *testing.T) {
	tm,err := time.Parse("02-01-2006", "23-11-2017")
	if err != nil {fmt.Println(err)}


	ti := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0,  time.UTC)
	if err != nil {fmt.Println(err)}
	//t.Fatalf("unexpected in out: ")
	out := []query{query{"",1511398800, 1511402400,ti,ti},
		query{"",1511402400, 1511406000,ti,ti},
		query{"",1511406000, 1511409600,ti,ti},
		query{"",1511409600, 1511413200,ti,ti},
		query{"",1511413200, 1511416800,ti,ti},
		query{"",1511416800, 1511420400,ti,ti},
		query{"",1511420400, 1511424000,ti,ti},
		query{"",1511424000, 1511427600,ti,ti},
		query{"",1511427600, 1511431200,ti,ti},
		query{"",1511431200, 1511434800,ti,ti},
		query{"",1511434800, 1511438000,ti,ti},
	}

	in := query{"",1511398800, 1511438000,ti,ti}
	if !testEq(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}

}


func TestGetDaysRanges_1(t *testing.T) {
	tm_1,err := time.Parse("02-01-2006", "23-11-2017")
	if err != nil {fmt.Println(err)}

	out := []query{query{"",0, 0, tm_1.AddDate(0,0,0),tm_1.AddDate(0,0,1)},
		query{"",0, 0,tm_1.AddDate(0,0,1),tm_1.AddDate(0,0,2)},
		query{"",0, 0,tm_1.AddDate(0,0,2),tm_1.AddDate(0,0,3)},
		query{"",0, 0,tm_1.AddDate(0,0,3),tm_1.AddDate(0,0,4)},
		query{"",0, 0,tm_1.AddDate(0,0,4),tm_1.AddDate(0,0,5)},
		query{"",0, 0,tm_1.AddDate(0,0,5),tm_1.AddDate(0,0,6)},
		query{"",0, 0,tm_1.AddDate(0,0,6),tm_1.AddDate(0,0,7)},
		query{"",0, 0,tm_1.AddDate(0,0,7),tm_1.AddDate(0,0,8)},
	}
	start,err :=time.Parse("02-01-2006", "23-11-2017")
	end,err :=time.Parse("02-01-2006", "30-11-2017")
	in := query{"",0, 0,start,end}
	if !testEq(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}

}

func testEq(a, b []query) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
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

	out := []RangeDate{RangeDate{1511398800, 1511402400,ti,ti},
		RangeDate{1511402400, 1511406000,ti,ti},
		RangeDate{1511406000, 1511409600,ti,ti},
		RangeDate{1511409600, 1511413200,ti,ti},
		RangeDate{1511413200, 1511416800,ti,ti},
		RangeDate{1511416800, 1511420400,ti,ti},
		RangeDate{1511420400, 1511424000,ti,ti},
		RangeDate{1511424000, 1511427600,ti,ti},
		RangeDate{1511427600, 1511431200,ti,ti},
		RangeDate{1511431200, 1511434800,ti,ti},
		RangeDate{1511434800, 1511438400,ti,ti},
	}

	in := RangeDate{1511398800, 1511438400,ti,ti}
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
	out := []RangeDate{RangeDate{1511398800, 1511402400,ti,ti},
		RangeDate{1511402400, 1511406000,ti,ti},
		RangeDate{1511406000, 1511409600,ti,ti},
		RangeDate{1511409600, 1511413200,ti,ti},
		RangeDate{1511413200, 1511416800,ti,ti},
		RangeDate{1511416800, 1511420400,ti,ti},
		RangeDate{1511420400, 1511424000,ti,ti},
		RangeDate{1511424000, 1511427600,ti,ti},
		RangeDate{1511427600, 1511431200,ti,ti},
		RangeDate{1511431200, 1511434800,ti,ti},
		RangeDate{1511434800, 1511438000,ti,ti},
	}

	in := RangeDate{1511398800, 1511438000,ti,ti}
	if !testEq(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}

}

func testEq(a, b []RangeDate) bool {

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
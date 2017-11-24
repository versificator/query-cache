package main

import (
	"testing"
)
func TestGetHourRanges_1(t *testing.T) {
	//t.Fatalf("unexpected in out: ")
	out := []rangeDate{rangeDate{1511398800, 1511402400},
		rangeDate{1511402400, 1511406000},
		rangeDate{1511406000, 1511409600},
		rangeDate{1511409600, 1511413200},
		rangeDate{1511413200, 1511416800},
		rangeDate{1511416800, 1511420400},
		rangeDate{1511420400, 1511424000},
		rangeDate{1511424000, 1511427600},
		rangeDate{1511427600, 1511431200},
		rangeDate{1511431200, 1511434800},
		rangeDate{1511434800, 1511438400},
	}

	in := rangeDate{1511398800, 1511438400}
	if !testEq(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}

}

func TestGetHourRanges_2(t *testing.T) {
	//t.Fatalf("unexpected in out: ")
	out := []rangeDate{rangeDate{1511398800, 1511402400},
		rangeDate{1511402400, 1511406000},
		rangeDate{1511406000, 1511409600},
		rangeDate{1511409600, 1511413200},
		rangeDate{1511413200, 1511416800},
		rangeDate{1511416800, 1511420400},
		rangeDate{1511420400, 1511424000},
		rangeDate{1511424000, 1511427600},
		rangeDate{1511427600, 1511431200},
		rangeDate{1511431200, 1511434800},
		rangeDate{1511434800, 1511438000},
	}

	in := rangeDate{1511398800, 1511438000}
	if !testEq(getHourRanges(in), out) {
		t.Fatalf("unexpected in out: %q; expecting: %q", getHourRanges(in), out)
	}

}

func testEq(a, b []rangeDate) bool {

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
package moons

import (
	"testing"
)

func TestCheckTideLocked(t *testing.T) {
	numMoons := []int{1, 3, 5}
	ranNum := []int{1, 6, 33, 89}
	results := [][][]bool{
		{[]bool{true, true}, []bool{true, false}, []bool{false, false}, []bool{false, false}},
		{[]bool{true, true}, []bool{true, false}, []bool{false, false}, []bool{false, false}},
		{[]bool{true, true}, []bool{true, true}, []bool{false, false}, []bool{false, false}},
	}
	for i, moons := range numMoons {
		for j, num := range ranNum {
			tide, big := CheckTideLocked(moons, num)
			result := results[i][j]
			if tide != result[0] || big != result[1] {
				t.Error("Expected", result[0], result[1], "received", tide, big)
			}
		}

	}

}

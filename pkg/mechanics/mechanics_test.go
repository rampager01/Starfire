package mechanics

import (
	"testing"
)

func TestFractionRoundUp(t *testing.T) {
	testNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testResults := [][]int{{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, {1, 1, 1, 2, 2, 2, 3, 3, 3, 4}}
	for i := 2; i < 4; i++ {
		for j, num := range testNums {
			x := FractionRoundUp(num, i)
			if x != testResults[i-2][j] {
				t.Error("Fraction Round Up expected", testResults[i-2][j], "but received", x)
			}
		}
	}
}

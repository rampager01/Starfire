package stars

import (
	"math/rand"
	"testing"
)

func BenchmarkCalculateTitusBode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CalculateTitusBode(BGStar, rand.NewSource(3))
	}
}

func TestCalculateTitusBode(t *testing.T) {
	for _, data := range Data {
		x := CalculateTitusBode(data.Star, rand.NewSource(3))
		if len(x) != len(data.TitusBode) {
			t.Error("Titus Bode Relationship not confirmed,\n got ", x, "\nExpected ", data.TitusBode)
		}
		for i, v := range x {
			if v != data.TitusBode[i] {
				t.Error("Titus Bode Relationship not confirmed,\n got ", x, "\nExpected ", data.TitusBode)
			}
		}
	}
}

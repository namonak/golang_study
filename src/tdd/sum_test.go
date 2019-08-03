package tdd

import (
	"math/rand"
	"testing"
)

func Test_SumInt32(t *testing.T) {
	loopCount := 1000
	
	for i := 0; i <= loopCount; i++ {
		a := []int32{rand.Int31n(10000), rand.Int31n(10000), rand.Int31n(10000), rand.Int31n(10000)}
		actual := SumInt32(a...)
		var expected int32
		for _, value := range a {
			expected += value
		}

		if actual != expected {
			t.Error("Invalid result ", actual, expected)
		}
	}
}

package tdd

import (
	"math/rand"
	"testing"
)

func Test_SumInt32(t *testing.T) {
	loopCount := 1000
	
	for i := 0; i <= loopCount; i++ {
		varArgCnt := rand.Intn(10)
		var a []int32
		for j := 0; j < varArgCnt; j++ {
			a = make([]int32, j+1)
		}
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

package bdd

import "math"

func getMaxInt32(n ...int32) int32 {
	var max int32 = math.MinInt32

	for _, value := range n {
		if value > max {
			max = value
		}
	}

	return max
}

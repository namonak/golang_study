package tdd

func SumInt32(n ...int32) int32 {
	var total int32

	for _, value := range n {
		total += value
	}	

	return total
}

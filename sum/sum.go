package sum

func sum(xs ...int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}

// func sum(a, b int) int {
// 	return a + b
// }
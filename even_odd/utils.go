package main

func arange(start int, stop int, size int) []int {
	rnge := make([]int, size)
	i := 0
	for x := start; x < stop; x++ {
		rnge[i] = x
		i += 1
	}
	return rnge
}

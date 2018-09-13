package data_structures

import (
	"fmt"
)
// a slice of slice ints
func main() {
	line1 := make([]int, 6, 6)
	line2 := make([]int, 6, 6)
	line3 := make([]int, 6, 6)
	line4 := make([]int, 6, 6)
	line5 := make([]int, 6, 6)
	line6 := make([]int, 6, 6)

	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &line1[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &line2[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &line3[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &line4[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &line5[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &line6[i])
	}

	//starting point
	/*
	first three nums
	next line down, one over from start point
	next line down, one over from start point
	starting point */

	counter := 0
	for i := 0; i < 100; i++ {

		for i := 0; i < 7; i++ {

		}
		counter += 1
	}
}

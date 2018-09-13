package hackerrank

import (
	"fmt"
	"math"
)

func main() {

	numofnums := Getindex()
	grid := CreateGrid(numofnums)
	finalnum := AddGrid(numofnums, grid)
	finalnumABS := math.Abs(float64(finalnum))
	fmt.Print(finalnumABS)

}

func Getindex() int {
	var index int
	fmt.Scanf("%d", &index)
	return index
}

func CreateGrid(numRowsClmns int) [][]int {
	grid := make([][]int, numRowsClmns)

	for y := 0; y < numRowsClmns; y++ {

		row := make([]int, numRowsClmns)
		for x := 0; x < numRowsClmns; x++ {
			fmt.Scanf("%d", &row[x])
		}
		grid[y] = row

	}
	return grid
}


func AddGrid(numRowsClmns int, grid [][]int) int {
	leftRight := 0
	for x, y := 0, 0; x < numRowsClmns; x, y = x+1, y+1 {
		leftRight += grid[y][x]
	}
	rightLeft := 0
	numRowsClmns = numRowsClmns - 1
	for x, y := numRowsClmns, 0; x >= 0; x, y = x-1, y+1 {
		rightLeft += grid[y][x]
	}
	finalNum := 0
	finalNum = leftRight-rightLeft
	return finalNum

}


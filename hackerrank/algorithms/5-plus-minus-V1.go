package hackerrank

import (
	"fmt"

)

func main() {

	index := getIndex()
	var numofPos, numofNegs, numofZeros int
	numofPos, numofNegs, numofZeros = adder(index)
	fractioner(index, numofPos, numofNegs, numofZeros)
}

func getIndex() int {
	var indexval int
	fmt.Scanf("%d", &indexval)
	return indexval
}

func adder(indexval int) (numofPos, numofNegs, numofZeros int) {
	currentNum := 0
	for i := 0; i < indexval; i++ {

		fmt.Scanf("%d", &currentNum)
		switch {
		case currentNum < 0:
			numofNegs += 1
		case currentNum > 0:
			numofPos += 1
		case currentNum == 0:
			numofZeros += 1


		}

	}
	return numofPos, numofNegs, numofZeros
}

func fractioner (indexval, numofPos, numofNegs, numofZeros int) {

	posRatio := float64(numofPos) / float64(indexval)
	negRatio := float64(numofNegs) / float64(indexval)
	zeroRatio := float64(numofZeros) / float64(indexval)
	fmt.Printf("%f\n", posRatio)
	fmt.Printf("%f\n", negRatio)
	fmt.Printf("%f\n", zeroRatio)
}
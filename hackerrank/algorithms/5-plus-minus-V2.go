package hackerrank

import (
	"fmt"

)

func main() {

	var indexval, numofPos, numofNegs, numofZeros, currentNum int
	fmt.Scanf("%d", &indexval)

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
	fmt.Println(float64(numofPos) / float64(indexval))
	fmt.Println(float64(numofNegs) / float64(indexval))
	fmt.Println(float64(numofZeros) / float64(indexval))

}


package hackerrank

import (
	"fmt"
	"math"
)


func main() {

	var index int
	fmt.Scanf("%d", &index)

	leftright := 0
	rightleft := 0
	leftrightT := 0
	rightleftT := 0
	for y := 0; y < index; y++ {

		for x := 0; x < index; x++ {
			dontcare := 0
			if x == y && x == index-y-1 {
				fmt.Scanf("%d", &leftright)
				leftrightT += leftright
				rightleft = leftright
				rightleftT += rightleft
			}else if x == y {
				fmt.Scanf("%d", &leftright)
				leftrightT += leftright
			}else if x == index-y-1 {
				fmt.Scanf("%d", &rightleft)
				rightleftT += rightleft
			}else {
				fmt.Scanf("%d", &dontcare)
			}


		}

	}
	fmt.Println(math.Abs(float64(leftrightT - rightleftT)))


}


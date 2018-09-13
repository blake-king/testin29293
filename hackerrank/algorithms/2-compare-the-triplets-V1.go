package hackerrank

import (
	"fmt"
)


func main() {

	var a0, b0, a1, b1, a2, b2 int
	fmt.Scanf("%d %d %d", &a0, &a1, &a2);
	fmt.Scanf("%d %d %d", &b0, &b1, &b2);
	var alpoints, bobpoints int

	switch {
	case a0 > b0:
		alpoints += 1
	case a0 < b0:
		bobpoints += 1
	}

	switch {
	case a1 > b1:
		alpoints += 1
	case a1 < b1:
		bobpoints += 1
	}

	switch {
	case a2 > b2:
		alpoints += 1
	case a2 < b2:
		bobpoints += 1
	}

	fmt.Print(alpoints,bobpoints)

}
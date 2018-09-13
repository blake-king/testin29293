package hackerrank

import (
	"fmt"
	"strings"
)

func main() {

	indexval := 0
	fmt.Scanf("%d", &indexval)
	for spaces := indexval-1; spaces >= 0; spaces-- {

		fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("#", indexval-spaces))
	}

}



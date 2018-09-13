package hackerrank

import (
"bytes"
"fmt"
)

func main() {

	indexval := 0
	fmt.Scanf("%d", &indexval)

	for i := 0; i < indexval; i++ {

		hts := i + 1
		spaces := indexval - hts

		var buffer bytes.Buffer

		for i := 0; i < spaces; i++ {

			buffer.WriteString(" ")
		}
		for i := 0; i < hts; i++ {
			buffer.WriteString("#")

		}

		fmt.Println(buffer.String())
	}

}


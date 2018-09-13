//

package hackerrank

import (
	"fmt"
)

func main() {
	testcases := 0
	fmt.Scanf("%d", &testcases)

	for i := 0; i < testcases; i++ {
		counter := 0
		numline := 0
		fmt.Scanf("%d", &numline)
		numslice := digits(numline)
		for i, _ := range numslice {

			if int(numslice[i]) == 0 {
				continue

			} else {
				if numline%int(numslice[i]) == 0 {
					counter++
				}
			}
		}
		fmt.Println(counter)
	}
}

func digits(dig int) []byte {

	digs := make([]byte, 0, 12)

	for {
		rest := dig / 10
		lastDig := dig - 10*rest
		digs = append(digs, byte(lastDig))
		if dig = rest; dig == 0 {
			break
		}
	}

	// reverse digits
	for i, j := 0, len(digs)-1; i < j; i, j = i+1, j-1 {
		digs[i], digs[j] = digs[j], digs[i]
	}

	return digs

}

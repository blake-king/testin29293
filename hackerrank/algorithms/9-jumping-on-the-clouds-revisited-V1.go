package hackerrank

import "fmt"


func main() {
	nvalue := 0
	kvalue := 0

	//clouds := [8]int{0, 0, 1, 0, 0, 1, 1, 0}
	energy := 100

	fmt.Scanf("%d", &nvalue)
	fmt.Scanf("%d", &kvalue)

	clouds := make([]int, nvalue)

	for i := 0; i < nvalue; i++ {

		fmt.Scanf("%d", &clouds[i])


	}


	idx := kvalue

	for i := 1; i <= nvalue/kvalue; i++ {
		if (kvalue * i) >= nvalue {
			idx = 0
		} else {
			idx = kvalue * i
		}

		energy = energy - 1
		if clouds[idx] == 1 {
			energy = energy - 2
		}

		/*fmt.Println(energy)*/
	}
	fmt.Print(energy)

}

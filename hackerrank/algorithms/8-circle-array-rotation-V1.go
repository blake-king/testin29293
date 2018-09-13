package hackerrank

import "fmt"


func main() {

	nvalue, kvalue, qvalue := 0, 0, 0

	fmt.Scanf("%d", &nvalue)
	fmt.Scanf("%d", &kvalue)
	fmt.Scanf("%d", &qvalue)

	arr := make([]int, nvalue, nvalue)
	arr2 := make([]int, nvalue, nvalue)
	qvals := make([]int, qvalue, qvalue)

	for i := 0; i < nvalue; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	rot := kvalue % nvalue
	newI := 0
	for i := 0; i < nvalue; i++ {

		if i >= nvalue-rot {
			arr2[newI] = arr[i]
			newI = newI + 1
			/*
			oldi := i
			for a := 0; a < rot; a, oldi = a+1, oldi+1 {
			arr2[a] = arr[oldi]


			}*/

		} else {
			arr2[i+rot] = arr[i]
		}


	}

	for i := 0; i < qvalue; i++ {
		fmt.Scanf("%d", &qvals[i])

	}

	for i := 0; i < qvalue; i++ {
		fmt.Println(arr2[qvals[i]])
	}

}

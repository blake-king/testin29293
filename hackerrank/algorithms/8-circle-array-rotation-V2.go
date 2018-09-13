package hackerrank

import "fmt"

/*n, how many numbers and the max size of those numbers less then (zero based index)

k - how many times rotates,

q - queries, each one has single integer of m

print element at index m in the rotated array
first line contains 3 ints,
n, k, q

n=3
k=2
q=3
second line contains n ints, where each int describes an array element

q number of more lines
contains single int denoting m*/
func main() {

	nvalue, kvalue, qvalue := 0, 0, 0

	fmt.Scanf("%d", &nvalue)
	fmt.Scanf("%d", &kvalue)
	fmt.Scanf("%d", &qvalue)

	arr := make([]int, nvalue, nvalue)
	arr2 := make([]int, nvalue, nvalue)
	qvals := make([]int, qvalue, qvalue)

	/*for i := 0; i < nvalue; i++ {
		fmt.Scanf("%d", &arr[i])
	}*/
	rot := kvalue % nvalue
	newI := 0
	for i := 0; i < nvalue; i++ {
		data := 0
		fmt.Scanf("%d", &data)
		if i >= nvalue-rot {
			arr2[newI] = data
			newI = newI + 1
			/*
			oldi := i
			for a := 0; a < rot; a, oldi = a+1, oldi+1 {
			arr2[a] = arr[oldi]


			}*/

		} else {
			arr2[i+rot] = data
		}


	}

	for i := 0; i < qvalue; i++ {
		fmt.Scanf("%d", &qvals[i])

	}

	for i := 0; i < qvalue; i++ {
		fmt.Println(arr2[qvals[i]])
	}

}

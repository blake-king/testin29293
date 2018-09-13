package hackerrank

import (
	"fmt"
	"testing"
	"io"
	"strings"
)

func newfunc(indexval int) {

	nvalue, kvalue, qvalue := 0, 0, 0

	var r io.Reader
	r = strings.NewReader("3 4 52 /")

	fmt.Scanf("%d", &nvalue)
	fmt.Scanf("%d", &kvalue)
	fmt.Scanf("%d", &qvalue)

	arr2 := make([]int, nvalue, nvalue)
	qvals := make([]int, qvalue, qvalue)


	rot := kvalue % nvalue
	newI := 0
	for i := 0; i < nvalue; i++ {
		data := 0
		fmt.Scanf("%d", &data)
		if i >= nvalue-rot {
			arr2[newI] = data
			newI = newI + 1

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

func oldfunc(indexval int) {

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

func Benchmarknewfunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		best(10)
	}
}

func Benchmarkoldfunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		better(10)
	}
}

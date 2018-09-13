package hackerrank

import "fmt"





func main () {
	numofnums := Assignindex()
	arr := make([]int64, numofnums)
	arr = Read(numofnums)
	finalresult := Add(arr, numofnums)
	fmt.Print(finalresult)
}

func Assignindex() int{
	var index int
	fmt.Scanf("%d" , &index)
	return index
}

func Read(index int) []int64{
	a := make([]int64, index)
	for i := range a {
		// ScanF Does work if you are assigning it to
		// an array. It separates them by spaces
		fmt.Scanf("%d", &a[i])
	}
	return a
}

func Add(arr []int64, index int) int64{
	var CurrentNum int64
	for i := 0; i < index; i++ {
		CurrentNum += arr[i]
	}
	return CurrentNum

}
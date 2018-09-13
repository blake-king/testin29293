package data_structures

import (
	"fmt"
	"strings"
)

func main(){


	var index int
	fmt.Scanf("%d", &index)
	arnums := make([]int, index)


	for i , j := 0, index-1; i < index; i, j = i+1, j-1 {
		fmt.Scanf("%d", &arnums[j])
	}

	fmt.Println(strings.Trim(fmt.Sprint(arnums), "[]"))

}
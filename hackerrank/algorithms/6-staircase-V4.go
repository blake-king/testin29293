package hackerrank


import (
	"bytes"
	"fmt"
)

func main(){

	indexval := 0
	fmt.Scanf("%d", &indexval)
	var buffer bytes.Buffer
	//create buffer
	var line = bytes.Repeat([]byte{0x20}, indexval+1) //space
	//line is all spaces at first, one extra char
	line[indexval] = 0x0a //newline
	//end of line = new line
	for i := indexval-1; i >= 0; i-- {
		line[i] = 0x23 //#
		buffer.Write(line)
	}
	fmt.Print(buffer.String())
}

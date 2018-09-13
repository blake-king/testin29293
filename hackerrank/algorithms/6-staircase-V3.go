package hackerrank


import (
	"fmt"
	"bytes"
)


func main() {

	indexval := 0
	fmt.Scanf("%d", &indexval)
	var line = bytes.Repeat([]byte{0x20}, indexval)

	var buffer bytes.Buffer
	for i := indexval-1; i >= 0; i-- {
		line[i] = 0x23
		buffer.Write(line)
		buffer.WriteByte(0x0a)
	}
	fmt.Print(buffer.String())
}
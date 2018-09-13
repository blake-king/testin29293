package hackerrank

import (
	"bytes"
	"strings"
	"testing"
)

//0x0a = newline
//0x20 = space
//0x23 = #


//flush?
//twice a smany function calls

func v4(indexval int) {
	var buffer bytes.Buffer
	//create buffer
	var line = bytes.Repeat([]byte{0x20}, indexval+1) //space
	//line is all spaces at first, one extra char
	line[indexval] = 0x0a //newline
	//end of line = new line
	for i := indexval - 1; i > 0; i-- {
		line[i] = 0x23 //#
		buffer.Write(line)
	}
	_ = buffer.String()
}

func v3(indexval int) {
	var line = bytes.Repeat([]byte{0x20}, indexval)
	var buffer bytes.Buffer
	for i := indexval - 1; i > 0; i-- {
		line[i] = 0x23
		buffer.Write(line)
		buffer.WriteByte(0x0a)
	}
	_ = buffer.String()
}

func v1(indexval int) {
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

		_ = buffer.String()
	}
}

func v2(indexval int) {
	result := ""
	for spaces := indexval - 1; spaces > 0; spaces-- {
		result = strings.Repeat(" ", spaces) + strings.Repeat("#", indexval-spaces) + "\n"
	}
	_ = result
}

func BenchmarkBest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		best(10)
	}
}

func BenchmarkBetter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v3(10)
	}
}

func BenchmarkOriginal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v1(10)
	}
}

func BenchmarkNotSoBad(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v2(10)
	}
}

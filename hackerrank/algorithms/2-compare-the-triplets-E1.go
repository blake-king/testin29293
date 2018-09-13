package hackerrank

import (
	"fmt"
	"strings"
	"strconv"
)


func main() {

	// Bad solution. No Scanln wouldnt work?


	var Alice, Bob string
	fmt.Scanln(&Alice)
	alres := strings.Fields(Alice)
	al1 := alres[0]
	al2 := alres[1]
	al3 := alres[2]
	al1int, _ := strconv.Atoi(al1)
	al2int, _ := strconv.Atoi(al2)
	al3int, _ := strconv.Atoi(al3)


	fmt.Scanln(&Bob)
	bobres := strings.Fields(Bob)
	bob1 := bobres[0]
	bob2 := bobres[1]
	bob3 := bobres[2]
	bob1int, _ := strconv.Atoi(bob1)
	bob2int, _ := strconv.Atoi(bob2)
	bob3int, _ := strconv.Atoi(bob3)

	var alpoints, bobpoints int

	switch  {
	case al1int > bob1int:
		alpoints += 1
	case al1int < bob1int:
		bobpoints += 1
	}

	switch  {
	case al2int > bob2int:
		alpoints += 1
	case al2int < bob2int:
		bobpoints += 1
	}

	switch  {
	case al3int > bob3int:
		alpoints += 1
	case al3int < bob3int:
		bobpoints += 1
	}


	fmt.Printf("%d %d", alpoints,bobpoints)

}
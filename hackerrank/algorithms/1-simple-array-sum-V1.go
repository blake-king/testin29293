package hackerrank
import "fmt"

func main() {
	var amntofNums, numberWorkingOn, sum int
	fmt.Scan(&amntofNums)
	for i:=0; i<amntofNums; i++ {
		fmt.Scan(&numberWorkingOn)
		sum += numberWorkingOn
	}
	fmt.Print(sum)
}
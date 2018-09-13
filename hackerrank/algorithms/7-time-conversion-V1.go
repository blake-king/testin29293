package hackerrank

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	data := ""
	fmt.Scanf("%s", &data)

	/*data := "01:05:45AM"*/
	parts := strings.Split(data, ":")

	if strings.Contains(parts[2], "PM") == true {

		if parts[0] != "12" {
			hour, _ := strconv.Atoi(parts[0])
			newHour := hour + 12
			newHourST := strconv.Itoa(newHour)
			parts[0] = newHourST
			parts[2] = strings.TrimSuffix(parts[2], "PM")
			fmt.Println(strings.Join(parts, ":"))
		} else {
			parts[2] = strings.TrimSuffix(parts[2], "PM")
			fmt.Println(strings.Join(parts, ":"))

		}
	} else {
		if parts[0] != "12" {
			parts[2] = strings.TrimSuffix(parts[2], "AM")
			fmt.Println(strings.Join(parts, ":"))
		} else {

			parts[2] = strings.TrimSuffix(parts[2], "AM")
			parts[0] = "00"
			fmt.Println(strings.Join(parts, ":"))
		}
	}

}

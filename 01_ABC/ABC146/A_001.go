package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	week := map[string]int{}
	week["MON"] = 0
	week["TUE"] = 1
	week["WED"] = 2
	week["THU"] = 3
	week["FRI"] = 4
	week["SAT"] = 5
	week["SUN"] = -1

	fmt.Println(6 - week[s])
}

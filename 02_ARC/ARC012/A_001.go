package main

import "fmt"

func main() {
	var day string
	fmt.Scan(&day)

	res := 0
	switch day {
	case "Monday":
		res = 5
	case "Tuesday":
		res = 4
	case "Wednesday":
		res = 3
	case "Thursday":
		res = 2
	case "Friday":
		res = 1
	}
	fmt.Println(res)
}

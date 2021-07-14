package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := len(s) - 1; i >= 0; i-- {
		var res string = s[i : i+1]
		if res == "6" {
			res = "9"
		} else if res == "9" {
			res = "6"
		}
		fmt.Print(res)
	}
	fmt.Println()
}

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s[len(s)-1] == 's' {
		fmt.Println(s + "es")
	} else {
		fmt.Println(s + "s")
	}
}

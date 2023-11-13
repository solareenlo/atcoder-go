package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if len(s) <= 2 {
		fmt.Println(0)
	} else {
		fmt.Println(s[:len(s)-2])
	}
}

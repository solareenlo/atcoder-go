package main

import "fmt"

func main() {
	var l int
	var s string
	fmt.Scan(&l, &s)

	if len(s) <= l {
		fmt.Println(s)
	} else {
		fmt.Println(s[:l])
	}
}

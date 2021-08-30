package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if len(a) > len(b) {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}

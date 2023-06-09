package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := strings.Split("xxxxxxxxxx", "")
	for i := 0; i < a; i++ {
		var x int
		fmt.Scan(&x)
		c[x] = "."
	}
	for i := 0; i < b; i++ {
		var x int
		fmt.Scan(&x)
		c[x] = "o"
	}
	fmt.Println(c[7], c[8], c[9], c[0])
	fmt.Println(c[4], c[5], c[6])
	fmt.Println(c[2], c[3])
	fmt.Println(c[1])
}

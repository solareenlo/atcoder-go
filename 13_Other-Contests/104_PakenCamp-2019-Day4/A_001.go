package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	c := N / 400
	if N%400 > 0 {
		c++
	}
	fmt.Println(c)
}

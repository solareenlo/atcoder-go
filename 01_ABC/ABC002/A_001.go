package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	res := x
	if x < y {
		res = y
	}
	fmt.Println(res)
}

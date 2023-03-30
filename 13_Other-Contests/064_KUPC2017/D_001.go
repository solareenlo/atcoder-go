package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := ""
	switch {
	case n <= 2:
		res = "1 1"
	case n == 3:
		res = "2 32"
	case n == 4:
		res = "2 18"
	default:
		res = "2 8"
	}
	fmt.Println(res)
}

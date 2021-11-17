package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res, x := 1, 1
	for i := 0; i < n; i++ {
		x <<= 1
		var s string
		fmt.Scan(&s)
		if s == "OR" {
			res += x
		}
	}

	fmt.Println(res)
}

package main

import "fmt"

func main() {
	var a, b, x, y int
	fmt.Scan(&a, &b, &x, &y)

	res := 0
	if a == b {
		res = x
	} else if a > b {
		res = min(x+(a-b-1)*y, x*((a-b)*2-1))
	} else if a < b {
		res = min(x+(b-a)*y, x*((b-a+1)*2-1))
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

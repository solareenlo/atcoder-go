package main

import "fmt"

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)
	res := 0
	if a >= k {
		res = 1
	} else if a <= b {
		res = -1
	} else {
		res = 1 + 2*((k-b-1)/(a-b))
	}
	fmt.Println(res)
}

package main

import "fmt"

func main() {
	var n, s, t int
	fmt.Scan(&n, &s, &t)
	res := 0
	for s < t {
		t /= 2
		res++
	}
	if s != t {
		res = -1
	}
	fmt.Println(res)
}

package main

import "fmt"

func main() {
	var t, a, b, c, d int
	fmt.Scan(&t, &a, &b, &c, &d)
	ans := 0
	if a <= t {
		ans = b
	}
	if c <= t {
		ans = d
	}
	if a+c <= t {
		ans = b + d
	}
	fmt.Println(ans)
}

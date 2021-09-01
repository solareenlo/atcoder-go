package main

import "fmt"

func main() {
	var a, b, c, k, s, t int
	fmt.Scan(&a, &b, &c, &k, &s, &t)

	if s+t >= k {
		fmt.Println((a-c)*s + (b-c)*t)
	} else {
		fmt.Println(a*s + b*t)
	}
}

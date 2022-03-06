package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	c := 1
	z := 0
	for i := 0; i < n; i++ {
		if i > 0 && s[i] == s[i-1] {
			z += c
			c++
		} else {
			c = 1
		}
	}
	fmt.Println(z)
}

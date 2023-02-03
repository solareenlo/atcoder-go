package main

import "fmt"

func main() {
	var x, k int
	fmt.Scan(&x, &k)

	c := 1
	for i := 1; i <= k; i++ {
		c *= 10
		m := x % 10
		if m >= 5 {
			x += 10
			x -= m
		} else {
			x -= m
		}
		x /= 10
	}
	fmt.Println(x * c)
}

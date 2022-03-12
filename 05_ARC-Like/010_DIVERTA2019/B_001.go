package main

import "fmt"

func main() {
	var a, b, c, n int
	fmt.Scan(&a, &b, &c, &n)

	s := 0
	for x := 0; x*a <= n; x++ {
		for y := 0; x*a+y*b <= n; y++ {
			if (n-x*a-y*b)%c < 1 {
				s++
			}
		}
	}
	fmt.Println(s)
}

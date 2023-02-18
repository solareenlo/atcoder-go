package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	a, b := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'X' {
			a++
		} else {
			b += a / 3
			a = 0
		}
	}
	fmt.Println(b)
}

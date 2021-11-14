package main

import "fmt"

func main() {
	var n, x int
	var s string
	fmt.Scan(&n, &x, &s)

	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			x++
		}
		if x != 0 && s[i] == 'x' {
			x--
		}
	}

	fmt.Println(x)
}

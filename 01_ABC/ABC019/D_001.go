package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var d, m, I int
	for i := 2; i <= n; i++ {
		fmt.Println("? 1", i)
		fmt.Scan(&d)
		if d > m {
			m, I = d, i
		}
	}

	m = 0
	for i := 1; i <= n; i++ {
		fmt.Println("?", I, i)
		fmt.Scan(&d)
		if d > m {
			m = d
		}
	}
	fmt.Println("!", m)
}

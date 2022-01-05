package main

import "fmt"

func main() {
	var n, m, a, b int
	fmt.Scan(&n, &m, &a, &b)

	ok := true
	day := 0
	for i := 0; i < m; i++ {
		if n <= a {
			n += b
		}
		var c int
		fmt.Scan(&c)
		n -= c
		if n < 0 {
			ok = false
			day = i + 1
			break
		}
	}

	if ok {
		fmt.Println("complete")
	} else {
		fmt.Println(day)
	}
}

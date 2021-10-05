package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n == 1 && m > 2 {
		fmt.Println(m - 2)
	} else if m == 1 && n > 2 {
		fmt.Println(n - 2)
	} else {
		fmt.Println((n - 2) * (m - 2))
	}
}

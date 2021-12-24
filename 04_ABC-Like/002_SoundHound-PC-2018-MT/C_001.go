package main

import "fmt"

func main() {
	var n, m, d float64
	fmt.Scan(&n, &m, &d)

	t := (n - d) * (m - 1.0) / n / n
	if d == 0.0 {
		fmt.Println(t)
	} else {
		fmt.Println(t * 2.0)
	}
}

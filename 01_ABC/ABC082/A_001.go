package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	mean := float64(a+b) / 2.0
	m := (a + b) / 2

	if mean-float64(m) > 0 {
		fmt.Println(m + 1)
	} else {
		fmt.Println(m)
	}
}

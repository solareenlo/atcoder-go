package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a, res := 0.0, 0.0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		res += 1.0 / a
	}
	fmt.Println(1.0 / res)
}

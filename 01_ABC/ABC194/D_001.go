package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0.0
	for i := 1; i < n; i++ {
		res += 1.0 / float64(i)
	}

	fmt.Println(res * float64(n))
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0.0
	for i := 0; i < n; i++ {
		res += float64(n) * 1.0 / (float64(n - i)) * 1.0
	}
	fmt.Println(res)
}

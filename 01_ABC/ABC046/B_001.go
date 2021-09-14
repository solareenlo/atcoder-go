package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	res := k
	for i := 0; i < n-1; i++ {
		res *= (k - 1)
	}
	fmt.Println(res)
}

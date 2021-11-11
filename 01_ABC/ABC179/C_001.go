package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	for i := 1; i < n; i++ {
		res += (n - 1) / i
	}

	fmt.Println(res)
}

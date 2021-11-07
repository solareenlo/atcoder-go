package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	for i := 1; i < n+1; i++ {
		res += (n / i) * (n/i + 1) * i / 2
	}

	fmt.Println(res)
}

package main

import "fmt"

func main() {
	const mod = 1000000007

	var N int
	fmt.Scan(&N)

	sum := 1
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		sum *= ((a + 2) % mod)
		sum %= mod
	}
	fmt.Println(sum - 2)
}

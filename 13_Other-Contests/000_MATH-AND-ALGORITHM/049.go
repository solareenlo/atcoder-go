package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	const mod = 1_000_000_007
	var fibo int
	fibo1 := 1
	fibo2 := 1
	for i := 3; i < n+1; i++ {
		fibo = (fibo1 + fibo2) % mod
		fibo1 = fibo2
		fibo2 = fibo
	}
	fmt.Println(fibo)
}

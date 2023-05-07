package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	i, c := 2, 0
	for ; i*i <= N; i++ {
		for ; N%i == 0; c++ {
			N /= i
		}
	}
	if N > 1 {
		c++
	}
	for N = 0; c > 1; c, N = c/2, N+1 {
		c++
	}
	fmt.Println(N)
}

package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var F [49]int
	F[0] = 1
	F[1] = 1
	for i := 2; i < 49; i++ {
		F[i] = F[i-1] + F[i-2]
	}
	ans := 0
	for i := 48; N > 0; i-- {
		ans += N / F[i]
		N %= F[i]
	}
	fmt.Println(ans)
}

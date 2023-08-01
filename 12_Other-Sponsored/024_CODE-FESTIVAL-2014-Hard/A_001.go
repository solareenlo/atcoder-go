package main

import "fmt"

func main() {
	var P float64
	var N int
	fmt.Scan(&P, &N)
	fmt.Println(0.5 * (1 - power(1-P*2, N)))
}

func power(a float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return power((a * a), n/2)
	}
	return (a * power(a, n-1))
}

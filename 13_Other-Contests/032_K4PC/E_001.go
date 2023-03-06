package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println("300 200")

	dp := make([]int, 201)
	dp[0] = 1
	var a, b int
	sum := 1
	for i := 1; i < 301; i++ {
		a = 0
		b = -1
		for j := 0; j < 200; j++ {
			a += dp[j]
			if sum+a <= n {
				b = j
			}
		}
		fmt.Println(200 - b)
		for j := 200; j >= 200-b; j-- {
			dp[j] += dp[j-200+b]
			sum += dp[j-200+b]
		}
	}
}

package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	sA := 0.0
	for i := 0; i < A; i++ {
		var a float64
		fmt.Scan(&a)
		sA += a
	}

	dp := make([]float64, 51)
	dp[0] = 1.0 / float64(B+C+1)
	for i := 0; i < B; i++ {
		var b float64
		fmt.Scan(&b)
		for j := i; j >= 0; j-- {
			dp[j+1] += dp[j] * b * float64(j+1) / float64(B+C-j)
		}
	}

	ans := 0.0
	for i := 0; i <= B; i++ {
		ans += sA * dp[i]
	}
	fmt.Println(ans)
}

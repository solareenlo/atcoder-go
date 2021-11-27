package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	L := [102]int{}
	R := [102]int{}
	for j := 0; j < T; j++ {
		dp := [102][102]int{}
		var N int
		fmt.Scan(&N)
		for i := 1; i < N+1; i++ {
			fmt.Scan(&L[i], &R[i])
			R[i]--
		}
		for len := 1; len < 101; len++ {
			for l, r := 1, len; r < 101; l, r = l+1, r+1 {
				bt := [102]int{}
				for i := 1; i < N+1; i++ {
					if l <= L[i] && R[i] <= r {
						bt[dp[l][L[i]-1]^dp[R[i]+1][r]] = 1
					}
				}
				for i := 0; i < 101; i++ {
					if bt[i] == 0 {
						dp[l][r] = i
						break
					}
				}
			}
		}
		if dp[1][100] != 0 {
			fmt.Println("Alice")
		} else {
			fmt.Println("Bob")
		}
	}
}

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	dp := [33][101010]int{}
	res := [101010]int{}

	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			dp[0][i] = i + 1
		} else {
			dp[0][i] = i - 1
		}
	}

	for j := 0; j < 32; j++ {
		for i := 0; i < n; i++ {
			dp[j+1][i] = dp[j][dp[j][i]]
		}
	}

	for i := 0; i < n; i++ {
		res[dp[32][i]]++
	}

	for i := 0; i < n; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(res[i])
	}
	fmt.Println()
}

package main

import "fmt"

func main() {
	var p, w int
	fmt.Scan(&p, &w)
	if p*w > 100 {
		fmt.Println(-2)
		return
	}
	if p*w == 100 {
		if w > 2 {
			fmt.Println(-2)
		} else {
			fmt.Println(-1)
		}
		return
	}
	if w <= 1 {
		fmt.Println(-1)
		return
	}
	ans := -1
	k := 0
	dp := make([]float64, 1)
	dp[0] = 1.0
	for k < ans+w+5 {
		k++
		dp = append(dp, 0.0)
		for j := len(dp) - 2; j >= 0; j-- {
			dp[j+1] = (float64(p)*dp[j] + (100.0-float64(p))*dp[j+1]) * 0.01
		}
		if dp[k/w+1] > 0.5 {
			ans = k
		}
	}
	fmt.Println(ans)
}

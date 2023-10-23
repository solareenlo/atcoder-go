package main

import "fmt"

func main() {
	var p [25]int = [25]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

	var a, b int
	fmt.Scan(&a, &b)

	var dp [3000]int
	dp[0] = 1

	for i := a; i <= b; i++ {
		x := 0
		for j := 0; j < 11; j++ {
			if i%p[j] == 0 {
				x |= (1 << j)
			}
		}
		for j := 0; j < (1 << 11); j++ {
			if (j & x) == 0 {
				dp[j|x] += dp[j]
			}
		}
	}

	ans := 0
	for i := 0; i < (1 << 11); i++ {
		ans += dp[i]
	}
	fmt.Println(ans)
}

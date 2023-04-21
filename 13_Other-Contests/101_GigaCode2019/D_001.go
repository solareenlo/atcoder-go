package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W, K, V int
	fmt.Fscan(in, &H, &W, &K, &V)

	var a [205][205]int
	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			fmt.Fscan(in, &a[i][j])
			a[i][j] += K
			a[i][j] += a[i-1][j]
		}
	}

	ans := 0
	for i := 1; i <= H; i++ {
		for j := i; j <= H; j++ {
			sum := 0
			for L, R := 1, 1; R <= W; R++ {
				sum += a[j][R] - a[i-1][R]
				for sum > V {
					sum -= a[j][L] - a[i-1][L]
					L++
				}
				ans = max(ans, (j-i+1)*(R-L+1))
			}
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

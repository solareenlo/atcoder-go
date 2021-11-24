package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	a := [2002][2002]int{}
	for i := 1; i < h+1; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 1; j < w+1; j++ {
			if s[j-1] == '+' {
				a[i][j] = 1
			} else {
				a[i][j] = -1
			}
		}
	}

	dp := [2002][2002]int{}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
	}
	dp[h][w] = 0

	for i := h; i > 0; i-- {
		for j := w; j > 0; j-- {
			if i == h && j == w {
				continue
			}
			dp[i][j] = max(a[i+1][j]-dp[i+1][j], a[i][j+1]-dp[i][j+1])
		}
	}

	if dp[1][1] > 0 {
		fmt.Println("Takahashi")
	} else if dp[1][1] < 0 {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Draw")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var mask [10000]int
	var dp [1 << 12]int
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < 12; j++ {
			if s[j] == 'o' {
				mask[i] |= 1 << j
			}
		}
	}

	for s := (1 << 12) - 1; s >= 0; s-- {
		for i := 0; i < n; i++ {
			if (s | mask[i]) > s {
				dp[s] = max(dp[s], dp[s|mask[i]]+1)
			}
		}
	}
	fmt.Println(dp[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

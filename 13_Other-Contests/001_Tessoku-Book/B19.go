package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, W int
	fmt.Fscan(in, &n, &W)
	dp := make([]int, 1000*n+1)
	for i := range dp {
		dp[i] = 1 << 30
	}
	dp[0] = 0
	s := 0
	for n > 0 {
		n--
		var w, v int
		fmt.Fscan(in, &w, &v)
		for j := s; j >= 0; j-- {
			dp[j+v] = min(dp[j+v], dp[j]+w)
		}
		s += v
	}
	for j := s; j >= 0; j-- {
		if dp[j] <= W {
			fmt.Println(j)
			return
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

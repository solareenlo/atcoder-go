package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var d, g int
	fmt.Scan(&d, &g)
	p := make([]int, d+1)
	c := make([]int, d+1)
	for i := 1; i <= d; i++ {
		fmt.Fscan(in, &p[i], &c[i])
	}
	dp := make([]int, 1011)
	var s int
	for i := 1; i <= d; i++ {
		s += p[i]
		for j := s; j > 0; j-- {
			for k := 1; k <= p[i] && k <= j; k++ {
				var tmp int
				if k == p[i] {
					tmp = c[i]
				}
				dp[j] = max(dp[j], dp[j-k]+k*100*i+tmp)
			}
		}
	}
	var i int
	for i = 0; dp[i] < g; i++ {
	}
	fmt.Println(i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

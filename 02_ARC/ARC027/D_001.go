package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	a -= b
	if a < 0 {
		a += mod
	}
	return a
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	h := make([]int, n)
	for i := range h {
		fmt.Fscan(in, &h[i])
	}

	var d int
	fmt.Fscan(in, &d)
	for i := 0; i < d; i++ {
		var s, t int
		fmt.Fscan(in, &s, &t)
		s--
		t--
		dp := make([]int, 16)
		imos := make([]int, 16)
		dp[s&15] = 1
		for i := s; i < t; i++ {
			d := i & 15
			dp[d] = add(dp[d], imos[d])
			imos[(i+1)&15] = add(imos[(i+1)&15], dp[d])
			imos[(i+h[i]+1)&15] = sub(imos[(i+h[i]+1)&15], dp[d])
			imos[(i+1)&15] = add(imos[(i+1)&15], imos[d])
			dp[d] = 0
			imos[d] = 0
		}
		dp[t&15] = add(dp[t&15], imos[t&15])
		fmt.Fprintln(out, dp[t&15])
	}
}

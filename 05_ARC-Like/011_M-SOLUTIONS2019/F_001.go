package main

import (
	"bufio"
	"fmt"
	"os"
)

func setAt(a []int, x int) {
	q := x / 64
	r := x % 64
	a[q] |= 1 << r
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	win := make([][]int, n)
	dp := make([][]int, n)
	len := (n + 63) / 64
	buf := make([]int, 2*len*n)
	for i := 0; i < n; i++ {
		win[i] = buf
		buf = buf[len:]
		dp[i] = buf
		buf = buf[len:]
		setAt(dp[i], i)
	}

	for i := 1; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < i; j++ {
			if s[j] == '1' {
				setAt(win[i], j)
			} else {
				setAt(win[j], i)
			}
		}
	}

	for len := 2; len <= n; len++ {
		for l := 0; l <= n-len; l++ {
			r := l + len - 1
			s := l / 64
			t := r/64 + 1
			for i := s; i < t; i++ {
				if win[l][i]&dp[l+1][i]&dp[r][i] != 0 {
					setAt(dp[r], l)
					break
				}
			}
			for i := s; i < t; i++ {
				if win[r][i]&dp[l][i]&dp[r-1][i] != 0 {
					setAt(dp[l], r)
					break
				}
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		q := i / 64
		r := i % 64
		ans += (dp[n-1][q] >> r) & (dp[0][q] >> r) & 1
	}
	fmt.Println(ans)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var m, S float64
	fmt.Fscan(in, &n, &m, &S)

	a := make([]float64, n+1)
	s := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		s[i] = s[i-1] + a[i]
	}

	ans := 0.0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if m*float64(n-j) >= S {
				ans = max(ans, (s[n]-s[j])*S/float64(n-j))
			} else {
				ans = max(ans, (s[n]-s[j])*m+(s[j]-s[i-1])*min(m, (S-float64(n-j)*m)/float64(j-i+1)))
			}
		}
	}
	fmt.Println(ans)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

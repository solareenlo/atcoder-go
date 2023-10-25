package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, S, dp_Bob, dp_Alice [1010]int

	var N int
	fmt.Fscan(in, &N)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i])
		S[i] = S[i-1] + A[i]
	}
	l, r := -int(1e18), int(1e18)
	for l+1 < r {
		m := (l + r) / 2
		for i := 1; i <= N; i++ {
			dp_Alice[i] = -1e9
			dp_Bob[i] = 1e9
		}
		dp_Alice[N+1] = 0
		dp_Bob[N+1] = 0
		for i := N; i >= 1; i-- {
			for j := N + 1; j >= i+1; j-- {
				tmp0 := -1
				if S[j-1]-S[i-1] >= m {
					tmp0 = 1
				}
				dp_Alice[i] = max(dp_Alice[i], dp_Bob[j]+tmp0)
				tmp1 := -1
				if S[j-1]-S[i-1] < m {
					tmp1 = 1
				}
				dp_Bob[i] = min(dp_Bob[i], dp_Alice[j]-tmp1)
			}
		}
		if dp_Alice[1] >= 0 {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(l)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

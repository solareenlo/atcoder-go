package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N  int
	L  int
	A  = make([]int, 2<<17)
	S  = make([]int, 2<<17)
	dp = make([]int, 2<<17)
)

func check(X int) int {
	for i := 0; i <= N; i++ {
		dp[i] = 0
	}
	dp[0] = 1
	dp[1] = -1
	l := 0
	r := 0
	for i := 0; i < N; i++ {
		dp[i+1] += dp[i]
		if dp[i] == 0 {
			continue
		}
		for l <= N && S[l]-S[i] < X {
			l++
		}
		for r <= N && S[r]-S[i] <= L {
			r++
		}
		dp[l]++
		dp[r]--
	}
	return dp[N]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &L)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		S[i+1] = S[i] + A[i]
	}
	l := 0
	r := L + 1
	for r-l > 1 {
		mid := (l + r) / 2
		if check(mid) != 0 {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Println(l)
}

package main

import "fmt"

var (
	N, X, Y int
	A       = [20]int{}
	B       = [20]int{}
	dp      = [1 << 18]int{}
)

func f(S, x int) int {
	ret := 0
	for p := 1; p <= N; p++ {
		if (S&(1<<(p-1))) == 0 && p < x {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Scan(&N, &X, &Y)
	for i := 1; i < N+1; i++ {
		fmt.Scan(&A[i])
	}
	for i := 1; i < N+1; i++ {
		fmt.Scan(&B[i])
	}

	dp[0] = 0
	for i := 1; i < 1<<N; i++ {
		dp[i] = 1 << 60
	}

	for S := 0; S < (1 << N); S++ {
		sizeS := 0
		for i := 1; i <= N; i++ {
			if S&(1<<(i-1)) != 0 {
				sizeS++
			}
		}
		for x := 1; x <= N; x++ {
			if S&(1<<(x-1)) != 0 {
				continue
			}
			dp[S|(1<<(x-1))] = min(dp[S|(1<<(x-1))], dp[S]+abs(A[x]-B[sizeS+1])*X+f(S, x)*Y)
		}
	}
	fmt.Println(dp[(1<<N)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, A, M, P, T, K int
	fmt.Fscan(in, &N, &A, &M, &P, &T, &K)
	var S string
	fmt.Fscan(in, &S)
	S = reverseString(S)
	powA := make([]int, N+1)
	for i := range powA {
		powA[i] = 1
	}
	for i := 0; i < N; i++ {
		powA[i+1] = (powA[i] * A) % P
	}
	var dp [101][601][601]int
	dp[0][0][0] = 1
	for i := 0; i < N; i++ {
		if S[i] == '0' {
			dp[i+1] = dp[i]
			continue
		}
		for sum := 0; sum < 1; sum++ {
			for rem := 0; rem < P; rem++ {
				n_sum := sum
				n_rem := rem
				tmp := 0
				for n_sum != M+1 {
					tmp += dp[i][n_sum][n_rem]
					if n_sum >= sum+A {
						tmp -= dp[i][n_sum-A][(P+(n_rem-A*powA[i])%P)%P]
					}
					dp[i+1][n_sum][n_rem] = min(K, tmp)
					n_sum++
					n_rem = (n_rem + powA[i]) % P
				}
			}
		}
	}
	for num := 1; num < K+1; num++ {
		if num > dp[N][M][T] {
			fmt.Fprintln(out, -1)
			return
		}
		tmp_num := num
		ans := make([]int, N)
		rem := T
		sum := M
		for i := N - 1; i >= 0; i-- {
			if S[i] == '0' {
				continue
			}
			for j := 0; j < A; j++ {
				n_rem := (rem + P - (powA[i]*j)%P) % P
				n_sum := (sum - j)
				if dp[i][n_sum][n_rem] < tmp_num {
					tmp_num -= dp[i][n_sum][n_rem]
				} else {
					ans[i] = j
					rem = n_rem
					sum = n_sum
					break
				}
			}
		}
		ans = reverseOrderInt(ans)
		for i := 0; i < N; i++ {
			if i != 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, ans[i])
		}
		fmt.Fprintln(out)
	}
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

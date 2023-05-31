package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353
	const size = 3000
	var nr [2 * (size + 1)][2 * (size + 1)]int
	nr[0][0] = 1
	for i := 0; i < size+1; i++ {
		for j := 0; j < size+1; j++ {
			nr[i][j+1] += nr[i][j]
			nr[i][j+1] %= mod
			nr[i+1][j] += nr[i][j]
			nr[i+1][j] %= mod
		}
	}
	var nCr func(int, int) int
	nCr = func(n, r int) int {
		if n < r || r < 0 {
			return 0
		}
		return nr[n-r][r]
	}

	var cnt1, cnt2 [6009]int
	var dp [6009][6009]int

	var N, M int
	fmt.Fscan(in, &N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &B[i])
	}
	sort.Ints(A)
	reverseOrderInt(A)
	sort.Ints(B)
	reverseOrderInt(B)
	for i := 1; i <= N+M; i++ {
		for j := 0; j < N; j++ {
			if A[j] < i {
				cnt1[i] += 1
			}
		}
	}
	for i := 0; i < M; i++ {
		cnt2[B[i]] += 1
	}

	dp[0][0] = 1
	for i := 1; i <= N+M; i++ {
		for j := 0; j < N+1; j++ {
			if dp[i-1][j] == 0 {
				continue
			}
			for k := 0; k <= cnt2[i]; k++ {
				if i == B[0] && B[0] > A[0] && k == 0 {
					continue
				}
				ways := nCr(cnt1[i]-j, k)
				dp[i][j+k] += dp[i-1][j] * ways
				dp[i][j+k] %= mod
			}
		}
	}

	ans := 0
	for i := 0; i < N+1; i++ {
		ans += dp[N+M][i]
		ans %= mod
	}
	fmt.Println(ans)
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func reverseOrderString(a []string) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD = 998244353

var N int
var dp []int

func divp(c int) {
	if c == 0 {
		return
	}
	for i := 0; i < N; i++ {
		if dp[i] != 0 {
			dp[i+1] = (dp[i+1] - dp[i]*c%MOD + MOD) % MOD
		}
	}
}

func mulp(c int) {
	if c == 0 {
		return
	}
	for i := N - 1; i >= 0; i-- {
		if dp[i] != 0 {
			dp[i+1] = (dp[i+1] + dp[i]*c%MOD) % MOD
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var fact, A, R, L, ret [2020]int
	fact[0] = 1
	for i := 0; i < 2010; i++ {
		fact[i+1] = fact[i] * (i + 1) % MOD
	}

	var Q int
	fmt.Fscan(in, &N, &Q)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	var ev [100][]pair
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &L[i], &R[i])
		ev[L[i]/45] = append(ev[L[i]/45], pair{R[i], i})
	}

	for i := 0; i < 46; i++ {
		if len(ev[i]) != 0 {
			CL := 45 * i
			CR := CL
			var cnt [2020]int
			dp = make([]int, 2020)
			dp[0] = 1
			sortPair(ev[i])
			for _, e := range ev[i] {
				for CR < e.x {
					divp(cnt[A[CR]])
					cnt[A[CR]]++
					mulp(cnt[A[CR]])
					CR++
				}
				for L[e.y] < CL {
					CL--
					divp(cnt[A[CL]])
					cnt[A[CL]]++
					mulp(cnt[A[CL]])
				}
				for CL < L[e.y] {
					divp(cnt[A[CL]])
					cnt[A[CL]]--
					mulp(cnt[A[CL]])
					CL++
				}
				for j := 0; j < N+1; j++ {
					if (j & 1) != 0 {
						ret[e.y] = (ret[e.y] - fact[N-j]*dp[j]) % MOD
					} else {
						ret[e.y] = (ret[e.y] + fact[N-j]*dp[j]) % MOD
					}
				}
			}
		}
	}

	for i := 0; i < Q; i++ {
		fmt.Fprintln(out, (ret[i]%MOD+MOD)%MOD)
	}
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

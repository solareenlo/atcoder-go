package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	type P struct {
		x, y int
	}

	var ps [2010][2010]int
	var buf, bug, p2 [2010]int
	p2[0] = 1
	for i := 0; i < 2010-1; i++ {
		p2[i+1] = p2[i] * 2 % MOD
	}

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for k := 1; k < N+1; k++ {
		for i := 0; i < k; i++ {
			buf[i] = r.Intn(MOD - 1)
		}
		for i := 0; i < N; i++ {
			if ((i / k) & 1) != 0 {
				ps[k][i+1] = (ps[k][i] + (((MOD - A[i]) % MOD) * buf[i%k] % MOD)) % MOD
			} else {
				ps[k][i+1] = (ps[k][i] + A[i]*buf[i%k]) % MOD
			}
		}
	}

	memo := make(map[P]int)

	for Q > 0 {
		Q--
		var l, r int
		fmt.Fscan(in, &l, &r)
		all0 := true
		for i := l; i <= r; i++ {
			if A[i] != 0 {
				all0 = false
			}
		}
		if all0 {
			fmt.Println(0)
			continue
		}
		if _, ok := memo[P{l, r}]; ok {
			fmt.Println(memo[P{l, r}])
			continue
		}
		n := r - l
		possibleks := make([]int, 0)
		for k := 1; k < n+1; k++ {
			if ps[k][r+1]-ps[k][l] == 0 {
				possibleks = append(possibleks, k)
			}
		}
		possibleks = reverseOrderINT(possibleks)
		ans := 0
		for i := 0; i < n+1; i++ {
			buf[i] = A[l+i]
		}
		var div func(int) bool
		div = func(k int) bool {
			if k > n {
				return false
			}
			for i := 0; i < n+1; i++ {
				bug[i] = buf[i]
			}
			for s := n; s > n-k; s-- {
				for i := s; i >= 0; i -= k {
					if i >= k {
						bug[i-k] = (bug[i-k] - bug[i] + MOD) % MOD
					} else if bug[i] != 0 {
						return false
					}
				}
			}
			for i := 0; i < n-k+1; i++ {
				buf[i] = bug[i+k]
			}
			n -= k
			return true
		}
		for _, k := range possibleks {
			for div(k) {
				ans = (ans + p2[k]) % MOD
			}
		}
		memo[P{l, r}] = ans
		fmt.Println(ans)
	}
}

func reverseOrderINT(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MOD = 1000000007
	const MX = 100009
	const B = 316

	var l, r, x, ans, sb, inv, frac1, frac2 [MX]int
	var prime [10009]int
	var s [70][MX]int
	var isprime [MX]bool
	for i := 1; i < 100000; i++ {
		isprime[i] = true
	}

	c := 0
	for i := 2; i <= 100000; i++ {
		if !isprime[i] {
			continue
		}
		for j := i * 2; j <= 100000; j += i {
			isprime[j] = false
		}
		prime[c] = i
		c++
	}
	inv[1] = 1
	for i := 2; i <= 100000; i++ {
		inv[i] = inv[MOD%i] * (MOD - MOD/i) % MOD
	}
	for i := 1; i <= 100000; i++ {
		frac1[i] = i * inv[i+1] % MOD
		frac2[i] = (i + 1) * inv[i] % MOD
	}

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i])
		sb[i] = i / B
		for j := 0; j < 65; j++ {
			for x[i]%prime[j] == 0 {
				x[i] /= prime[j]
				s[j][i+1]++
			}
		}
		if x[i] == 1 {
			x[i] = -1
		} else {
			x[i] = lowerBound(prime[:c], x[i])
		}
	}
	for i := 0; i < 65; i++ {
		for j := 0; j < N; j++ {
			s[i][j+1] += s[i][j]
		}
	}
	p := make([]int, Q+1)
	for i := 1; i <= Q; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		l[i]--
		p[i] = i
	}
	sort.Slice(p, func(i, j int) bool {
		if sb[l[p[i]]] != sb[l[p[j]]] {
			return sb[l[p[i]]] < sb[l[p[j]]]
		}
		if (sb[l[p[i]]] & 1) != 0 {
			return r[p[i]] > r[p[j]]
		}
		return r[p[i]] < r[p[j]]
	})

	cnt := make([]int, c)
	for i := range cnt {
		cnt[i] = 1
	}
	ret, pl, pr := 1, 0, 0
	for i := 1; i <= Q; i++ {
		for pl > l[p[i]] {
			pl--
			if x[pl] != -1 {
				ret = ret * frac2[cnt[x[pl]]] % MOD
				cnt[x[pl]]++
			}
		}
		for pr < r[p[i]] {
			if x[pr] != -1 {
				ret = ret * frac2[cnt[x[pr]]] % MOD
				cnt[x[pr]]++
			}
			pr++
		}
		for pl < l[p[i]] {
			if x[pl] != -1 {
				cnt[x[pl]]--
				ret = ret * frac1[cnt[x[pl]]] % MOD
			}
			pl++
		}
		for pr > r[p[i]] {
			pr--
			if x[pr] != -1 {
				cnt[x[pr]]--
				ret = ret * frac1[cnt[x[pr]]] % MOD
			}
		}
		res := ret
		for j := 0; j < 65; j++ {
			res = res * (s[j][r[p[i]]] - s[j][l[p[i]]] + 1) % MOD
		}
		ans[p[i]] = res
	}
	for i := 1; i <= Q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

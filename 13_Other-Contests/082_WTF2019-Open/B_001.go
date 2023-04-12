package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const Maxn = 40
const Maxs = 1 << 17

type Q struct {
	x, y int
}

var n, q, tot int
var a, lim [Maxn]Q
var c, A, B, id, fa [Maxn]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &q)
	for i := 1; i <= q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[i] = Q{x - 1, y}
		tot++
		c[tot] = x - 1
		tot++
		c[tot] = y
	}
	tmp := c[1 : tot+1]
	sort.Ints(tmp)
	tot = len(unique(c[1 : tot+1]))
	for i := 1; i <= tot; i++ {
		fa[i] = i
	}
	for i := 1; i <= q; i++ {
		x := lowerBound(c[1:tot+1], a[i].x) + 1
		y := lowerBound(c[1:tot+1], a[i].y) + 1
		merge(x, y)
	}
	cnt := 0
	for i := 1; i <= tot; i++ {
		fx := find(i)
		if fx == i {
			cnt++
			id[i] = cnt
		}
	}
	for i := 2; i <= tot; i++ {
		A[i-1] = powMod(9, mod-2) * (powMod(10, c[i]-c[i-1]) + mod - 1) % mod
		B[i-1] = (powMod(A[i-1], mod-2) + 1) % mod
		x := id[find(i-1)]
		y := id[find(i)]
		lim[i-1] = Q{x, y}
	}
	if q == 0 {
		fmt.Print(powMod(10, n))
	}
	fmt.Println(work(cnt, tot-1) * powMod(10, n-c[tot]+c[1]) % mod)
}

func find(x int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func merge(x, y int) {
	fx := find(x)
	fy := find(y)
	if fx != fy {
		fa[fx] = fy
	}
}

var dp, f, val [Maxs]int

func work(n, m int) int {
	for i := 1; i < (1 << n); i++ {
		val[i] = 1
		for j := 1; j <= m; j++ {
			if (i&(1<<(lim[j].x-1))) != 0 && (i&(1<<(lim[j].y-1))) != 0 {
				val[i] = val[i] * B[j] % mod
			}
		}
	}
	for i := 1; i < (1 << n); i++ {
		f[i] = val[i]
		st := i - lowbit(i)
		for j := st; j > 0; j = (j - 1) & st {
			f[i] = (f[i] - f[i^j]*val[j]%mod + mod) % mod
		}
	}
	for i := 0; i < (1 << n); i++ {
		f[i] *= 9
		f[i] %= mod
	}
	dp[0] = 1
	for i := 1; i < (1 << n); i++ {
		dp[i] = f[i]
		st := i - lowbit(i)
		for j := st; j > 0; j = (j - 1) & st {
			dp[i] += dp[j] * f[i^j] % mod
			dp[i] %= mod
		}
	}
	ans := powMod(9, mod-2) * dp[(1<<n)-1] % mod
	for i := 1; i <= m; i++ {
		ans = ans * A[i] % mod
	}
	return ans
}

func lowbit(x int) int {
	return x & -x
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353
const N = 200005

var n, sts int
var p, f, g, st [N]int
var tr [N*4 + 5]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	build(1, n, 1)
	f[0] = 1
	g[0] = 1
	for i := 1; i <= n; i++ {
		for sts != 0 && p[i] < p[st[sts]] {
			sts--
		}
		if sts == 0 {
			f[i] = g[i-1]
		} else {
			f[i] = (g[i-1] - g[st[sts]-1] + MOD) % MOD
			pos := query(1, n, st[sts], p[i], 1)
			var tp int
			if pos != 0 {
				tp = find(pos)
				f[i] = (f[i] + f[tp]) % MOD
				f[i] = (f[i] - (g[tp-1]-g[pos-1]+MOD)%MOD + MOD) % MOD
			}
		}
		g[i] = (g[i-1] + f[i]) % MOD
		sts++
		st[sts] = i
	}
	fmt.Println(f[n])
}

func build(l, r, id int) {
	if l == r {
		tr[id] = p[l]
		return
	}
	mid := (l + r) >> 1
	build(l, mid, id<<1)
	build(mid+1, r, id<<1|1)
	tr[id] = max(tr[id<<1], tr[id<<1|1])
}

func query(l, r, en, v, id int) int {
	if tr[id] < v {
		return 0
	}
	if l == r {
		return l
	}
	mid := (l + r) >> 1
	ans := 0
	if mid < en {
		ans = query(mid+1, r, en, v, id<<1|1)
	}
	if ans == 0 {
		ans = query(l, mid, en, v, id<<1)
	}
	return ans
}

func find(x int) int {
	l := 1
	r := sts
	var mid int
	for l <= r {
		mid = (l + r) >> 1
		if st[mid] >= x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return st[r+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

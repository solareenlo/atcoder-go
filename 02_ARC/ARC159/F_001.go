package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 1000010
const mod = 998244353

var a, dp, c, S, st [maxn]int
var vis [maxn]bool
var top, lim int

func chk(x int) int {
	if x > lim {
		return lim + 1
	}
	if x < 0 {
		return 0
	}
	return x
}

func cdq(l, r int) {
	if l == r {
		if a[l+l-1] != a[l+l] {
			dp[l] = (dp[l] + dp[l-1]) % mod
		}
		return
	}
	mid := (l + r) >> 1
	cdq(l, mid)
	Sum := 0
	top = 0
	for i := l; i <= mid; i++ {
		Sum = (Sum + dp[i-1]) % mod
	}
	for i := mid + 1; i <= r; i++ {
		dp[i] = (dp[i] + Sum) % mod
	}
	Mx, id := 0, 0
	for i := mid + 1; i <= r; i++ {
		ls := a[i+i-1]
		rs := a[i+i]
		c[ls]++
		if Mx < c[ls] {
			Mx = c[ls]
			id = ls
		}
		c[rs]++
		if Mx < c[rs] {
			Mx = c[rs]
			id = rs
		}
		if Mx > i-mid && !vis[id] {
			top++
			st[top] = id
			vis[id] = true
		}
	}
	Mx = 0
	for i := mid + 1; i <= r; i++ {
		c[a[i+i-1]] = 0
		c[a[i+i]] = 0
	}
	for i := mid; i >= l; i-- {
		ls := a[i+i-1]
		rs := a[i+i]
		c[ls]++
		if Mx < c[ls] {
			Mx = c[ls]
			id = ls
		}
		c[rs]++
		if Mx < c[rs] {
			Mx = c[rs]
			id = rs
		}
		if Mx > mid-i+1 && !vis[id] {
			top++
			st[top] = id
			vis[id] = true
		}
	}
	for i := l; i <= mid; i++ {
		c[a[i+i-1]] = 0
		c[a[i+i]] = 0
	}
	for i := 1; i <= top; i++ {
		vis[st[i]] = false
	}
	lim = 4 * (mid - l + 1)
	bas := lim / 2
	S[lim+1] = 0
	for o := 1; o <= top; o++ {
		v := st[o]
		cur := bas
		for i := 0; i <= lim; i++ {
			S[i] = 0
		}
		for i := mid; i >= l; i-- {
			ls := a[i+i-1]
			rs := a[i+i]
			tmp0 := -1
			if ls == v {
				tmp0 = 1
			}
			tmp1 := -1
			if rs == v {
				tmp1 = 1
			}
			cur += tmp0 + tmp1
			S[cur] = (S[cur] + dp[i-1]) % mod
		}
		cur = bas
		for i := lim; i >= 0; i-- {
			S[i] = (S[i] + S[i+1]) % mod
		}
		for i := mid + 1; i <= r; i++ {
			ls := a[i+i-1]
			rs := a[i+i]
			tmp0 := -1
			if ls == v {
				tmp0 = 1
			}
			tmp1 := -1
			if rs == v {
				tmp1 = 1
			}
			cur += tmp0 + tmp1
			dp[i] = (dp[i] - S[chk(lim+1-cur)] + mod) % mod
		}
	}
	cdq(mid+1, r)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	dp[0] = 1
	for i := 1; i <= n*2; i++ {
		fmt.Fscan(in, &a[i])
	}
	cdq(1, n)
	fmt.Println(dp[n])
}

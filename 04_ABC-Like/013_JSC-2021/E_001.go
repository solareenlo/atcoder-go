package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 500005 + 5

var (
	s   = [M]byte{}
	ga  int
	c   = [M]int{}
	idx int
	dp  = [M][30]int{}
	g   = [M]int{}
	g1  = [M]int{}
)

func dfs(l, r, d int) {
	if d == 0 {
		for i := l; i <= r; i++ {
			idx++
			c[i] = idx
		}
		return
	}
	l1, r1, r2 := l, 0, r
	if (l+r)&1 == 0 {
		r1 = (l+r)/2 - 1
	} else {
		r1 = (l + r) / 2
	}
	dfs(l1, r1, d-1)
	for i := 0; i < r1-l1+1; i++ {
		c[r2-i] = c[i+l1]
	}
	if (l+r)%2 == 0 {
		idx++
		c[(l+r)/2] = idx
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var ga int
	var s string
	fmt.Fscan(in, &ga, &s)
	n := len(s)
	s = "#" + s
	p := n
	for i := 1; i <= ga; i++ {
		if p == 0 {
			fmt.Println("impossible")
			return
		}
		p >>= 1
	}
	if p == 1 {
		fmt.Println("impossible")
		return
	}
	dfs(1, n, ga)
	for i := 1; i <= n; i++ {
		dp[c[i]][s[i]-'a'+1]++
	}
	for i := 1; i <= idx; i++ {
		for j := 1; j <= 26; j++ {
			if dp[i][j] > dp[i][g[i]] {
				g1[i] = g[i]
				g[i] = j
			} else if dp[i][j] > dp[i][g1[i]] {
				g1[i] = j
			}
		}
	}
	flag := true
	for i := 1; i <= p/2; i++ {
		if g[i] != g[p-i+1] {
			flag = false
			break
		}
	}
	if flag && p > 0 {
		t := 1
		for i := 1; i <= p; i++ {
			if i+i-1 != p && dp[i][g[i]]-dp[i][g1[i]] < dp[t][g[t]]-dp[t][g1[t]] {
				t = i
			}
		}
		g[t] = g1[t]
	}
	res := 0
	for i := 1; i <= idx; i++ {
		res += dp[i][g[i]]
	}
	fmt.Println(n - res)
}

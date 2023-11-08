package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200200

var t [N][26]int
var tot int = 0
var s string
var vt [N][]int
var n int

func ins(n, id int) {
	x := 1
	for i := 1; i <= n; i++ {
		if t[x][s[i]-'a'] != 0 {
			x = t[x][s[i]-'a']
		} else {
			tot++
			t[x][s[i]-'a'] = tot
			x = tot
		}
	}
	vt[x] = append(vt[x], id)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nxt, c, fir, ans [N]int
	var vis [N * 20]bool

	fmt.Fscan(in, &n)
	tot = 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s)
		m := len(s)
		s = " " + s
		for i := 2; i <= m; i++ {
			k := nxt[i-1]
			for k > 0 && s[k+1] != s[i] {
				k = nxt[k]
			}
			if s[k+1] == s[i] {
				k++
			}
			nxt[i] = k
		}
		b := 1
		if m%(m-nxt[m]) == 0 {
			b = m / (m - nxt[m])
			m -= nxt[m]
		}
		c[i] = b
		ins(m, i)
	}
	mx := 0
	for i := 1; i <= n; i++ {
		mx = max(mx, c[i])
	}
	for i := 1; i <= mx; i++ {
		fir[i] = i
	}
	for i := 1; i <= tot; i++ {
		for _, j := range vt[i] {
			x := c[j]
			for vis[fir[x]] {
				fir[x] += x
			}
			vis[fir[x]] = true
			ans[j] = fir[x] / x
		}
		for _, j := range vt[i] {
			x := c[j]
			fir[x] = x
			vis[ans[j]*x] = false
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

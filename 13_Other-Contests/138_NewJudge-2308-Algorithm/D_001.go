package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 300100
const M = 600100
const INF = int(1e18)

var t, ans int
var h, q, s [N]int
var f [N]bool
var w, l [M]int

func push(x, y int) {
	t++
	w[t] = y
	l[t] = h[x]
	h[x] = t
}

func bfs(x int) int {
	ans = -INF
	q[1] = x
	s[1] = 1
	f[x] = true
	for i, j := 1, 1; i <= j; i++ {
		ans = max(ans, s[i])
		for k := h[q[i]]; k > 0; k = l[k] {
			if !f[w[k]] {
				f[w[k]] = true
				j++
				q[j] = w[k]
				s[j] = s[i] + 1
			}
		}
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n1, n2, m int
	fmt.Fscan(in, &n1, &n2, &m)
	for m > 0 {
		m--
		var a, b int
		fmt.Fscan(in, &a, &b)
		push(a, b)
		push(b, a)
	}
	fmt.Println(bfs(1) + bfs(n1+n2) - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

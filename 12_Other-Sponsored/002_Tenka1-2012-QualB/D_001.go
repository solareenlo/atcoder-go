package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const INF = 500000000

var h, w int
var buf [20]string
var g [40][]int
var match [40]int
var vis [40]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &h, &w)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &buf[i])
	}
	fmt.Println(solve())
}

func solve() int {
	empty := true
	full := true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if buf[i][j] == '#' {
				empty = false
			} else {
				full = false
			}
		}
	}
	if empty {
		return 0
	}
	if full {
		return -1
	}

	res := INF
	for bit := 0; bit < (1 << h); bit++ {
		cnt := popcount(uint32(bit))
		xbit := 0
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if buf[i][j] == '#' && ((bit>>i)&1) == 0 {
					xbit |= (1 << j)
				}
			}
		}
		xcnt := popcount(uint32(xbit))
		for i := 0; i < h+w; i++ {
			g[i] = make([]int, 0)
		}

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if (bit>>i)&1 != 0 && (xbit>>j)&1 != 0 && buf[i][j] == '.' {
					g[i] = append(g[i], h+j)
					g[h+j] = append(g[h+j], i)
				}
			}
		}
		matching := 0
		for i := range match {
			match[i] = -1
		}
		for i := 0; i < h; i++ {
			if match[i] == -1 {
				for j := 0; j < h+w; j++ {
					vis[j] = false
				}
				if dfs(i) {
					matching++
				}
			}
		}
		if matching == 0 {
			matching = -1
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if buf[i][j] == '.' && (((bit>>i)&1) != 0 || ((xbit>>j)&1) != 0) {
						matching = 0
					}
				}
			}
		}
		res = min(res, xcnt+cnt-matching)
	}
	return res
}

func dfs(v int) bool {
	vis[v] = true
	for i := 0; i < len(g[v]); i++ {
		to := g[v][i]
		w := match[to]
		if w == -1 || (!vis[w] && dfs(w)) {
			match[v] = to
			match[to] = v
			return true
		}
	}
	return false
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

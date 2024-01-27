package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var n int

func calc(u, w int) int {
	if w >= 61 {
		return 0
	}
	if n/u < (1 << w) {
		return 0
	}
	return min(n-(u<<w)+1, 1<<w)
}

func solve() {
	var x, k int
	fmt.Fscan(in, &n, &x, &k)
	if k == 0 {
		fmt.Println(1)
		return
	}
	ans := 0
	ans += calc(x, k)
	if x>>min(k, 61) != 0 {
		ans++
	}
	u := x >> 1
	uu := x
	for u != 0 && k >= 2 {
		s := u * 2
		if s == uu {
			s = u*2 + 1
		}
		if k >= 2 {
			ans += calc(s, k-2)
		}
		k--
		uu = u
		u >>= 1
	}
	fmt.Println(ans)
}

func main() {

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		solve()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

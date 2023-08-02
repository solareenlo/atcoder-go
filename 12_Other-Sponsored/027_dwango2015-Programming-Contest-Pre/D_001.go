package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n, L int
	fmt.Fscan(in, &n, &L)
	var x, d [1 << 17]int
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &x[i], &a, &b)
		d[i] = a - b
	}
	c, w := 0, 0
	sum := 0
	t := make([]pair, n)
	for i := 0; i < n; i++ {
		b := x[(i+1)%n] - x[i]
		if b < 0 {
			b += L
		}
		c += d[i]
		sum += b * c
		t[i] = pair{c, b}
		w += b
	}
	sortPair(t)
	ans := INF
	for i := 0; i < n; i++ {
		ans = min(ans, sum-w*t[i].x)
		w -= 2 * t[i].y
		sum -= 2 * t[i].y * t[i].x
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

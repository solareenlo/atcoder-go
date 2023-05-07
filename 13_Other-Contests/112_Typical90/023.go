package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	type pair struct {
		x, y int
	}

	var h, w int
	fmt.Fscan(in, &h, &w)

	var t []int
	for i := 0; i < 1<<(w+2); i++ {
		if (i & (i >> 1)) == 0 {
			t = append(t, i)
		}
	}

	g := make([]pair, 0)
	for i := 0; i < len(t); i++ {
		x := lowerBound(t, t[i]>>1)
		y := -1
		if (t[i]&7 | t[i]>>(w+1)) == 0 {
			y = lowerBound(t, t[i]>>1|1<<(w+1))
		}
		g = append(g, pair{x, y})
	}

	d := make([]int, len(t))
	d[0] = 1
	for i := 0; i < h+1; i++ {
		c := strings.Repeat("#", w)
		if i < h {
			fmt.Fscan(in, &c)
		}
		c += "#"
		for _, a := range c {
			e := make([]int, len(t))
			for j := 0; j < len(t); j++ {
				x := g[j].x
				y := g[j].y
				e[x] = (e[x] + d[j]) % MOD
				if y >= 0 && y < len(e) && a != '#' {
					e[y] = (e[y] + d[j]) % MOD
				}
			}
			d, e = e, d
		}
	}
	fmt.Println(d[0])
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

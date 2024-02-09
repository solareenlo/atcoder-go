package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	ans := 0
	d := make([]pair, 0)
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if c == 0 {
			d = append(d, pair{a, b - 1})
		} else {
			if b == 1 {
				d = append(d, pair{a, 1})
				continue
			}
			d = append(d, pair{2 * a, 1})
			d = append(d, pair{a, b - 2})
		}
	}
	sortPair(d)
	for i := range d {
		u, v := d[i].x, d[i].y
		tmp := min(m, v)
		ans += u * (v - tmp)
		m -= tmp
	}
	fmt.Println(ans)
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y > tmp[j].y
		}
		return tmp[i].x > tmp[j].x
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

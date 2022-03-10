package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)

	e := [1005][1005]int{}
	m := 0
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		e[l%d][r%d]++
		m = max(m, e[l%d][r%d])
	}

	var s int
	for s = 1; s*s < m; s++ {
	}

	h := make([]bool, d)
	w := make([]bool, d)
	bt := [1005][1005]bool{}
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			if s*(s-1) < e[i][j] {
				h[i] = true
				w[j] = true
			} else if (s-1)*(s-1) < e[i][j] {
				bt[i][j] = true
			}
		}
	}

	l := 0
	r := d
	for l+1 < r {
		c := [1005]int{}
		o := 0
		ok := false
		m = (l + r) / 2
		for i := 0; i < d; i++ {
			for ; o < i+m && !h[o%d]; o++ {
				for j := 0; j < d; j++ {
					if bt[o%d][j] {
						c[j]++
					}
				}
			}
			if o < i+m {
				for j := 0; j < d; j++ {
					c[j] = 0
				}
				i = o
				o++
				continue
			}
			p := 0
			for j := 0; j < d; j++ {
				for ; p < j+m && !w[p%d] && c[p%d] == 0; p++ {

				}
				if p < j+m {
					j = p
					p++
				} else {
					ok = true
				}
			}
			for j := 0; j < d; j++ {
				if bt[i][j] {
					c[j]--
				}
			}
		}
		if ok {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(s*d - l - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

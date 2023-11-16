package main

import (
	"fmt"
)

type P struct {
	x, y int
}

var h, w int
var cnta, cntb [2][100100]int
var cntc, cntd [2][225816]int
var st map[P]bool
var ans int = 0

func check(i, j int) bool {
	return (0 <= i && i < h && 0 <= j && j < w)
}

func solve(k, r, c int) {
	cnta[k][r]++
	cntb[k][c]++
	cntc[k][r+c]++
	cntd[k][r+w-c]++
	if k != 0 {
		if _, ok := st[P{r, c}]; ok {
			ans -= 4
		}
	} else {
		st[P{r, c}] = true
	}
}

func main() {

	st = make(map[P]bool)

	var x, y [2]int

	fmt.Scan(&h, &w)
	n := h + w
	for k := 0; k < 2; k++ {
		fmt.Scan(&x[k], &y[k])
		x[k]--
		y[k]--
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if !(i != 0 || j != 0) {
					continue
				}
				s := x[k]
				t := y[k]
				for {
					s += i
					t += j
					if !check(s, t) {
						break
					}
					solve(k, s, t)
				}
			}
		}
	}
	for i := 0; i < h; i++ {
		ans += cnta[0][i] * cnta[1][i]
	}
	for i := 0; i < w; i++ {
		ans += cntb[0][i] * cntb[1][i]
	}
	for i := 0; i < n; i++ {
		ans += cntc[0][i] * cntc[1][i]
		ans += cntd[0][i] * cntd[1][i]
	}
	fmt.Println(ans)
}

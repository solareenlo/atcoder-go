package main

import (
	"fmt"
)

func main() {
	var hh, ww int
	fmt.Scan(&hh, &ww)
	h := max(ww, hh)
	w := min(ww, hh)
	rev := hh < ww

	g := make([][]int, h)
	for i := range g {
		g[i] = make([]int, w)
	}
	if !rev {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				fmt.Scan(&g[i][j])
			}
		}
	} else {
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				fmt.Scan(&g[j][i])
			}
		}
	}

	var n int
	fmt.Scan(&n)
	v := make([][]int, h)
	for i := range v {
		v[i] = make([]int, w)
		for j := range v[i] {
			v[i][j] = 1
		}
	}
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if rev {
			x, y = y, x
		}
		v[y][x] = 0
	}

	ans := 100000
	for f := 0; f < (1 << w); f++ {
		ok := true
		for i := 0; i < w; i++ {
			if v[0][i] == 0 && (f&(1<<i)) != 0 {
				ok = false
			}
		}
		if !ok {
			continue
		}
		d := make([][]int, h)
		for i := range d {
			d[i] = make([]int, w)
			copy(d[i], g[i])
		}
		cnt := 0
		push := func(x, y int) {
			d[y][x] = 1 - d[y][x]
			if y != 0 {
				d[y-1][x] = 1 - d[y-1][x]
			}
			if y != h-1 {
				d[y+1][x] = 1 - d[y+1][x]
			}
			if x != 0 {
				d[y][x-1] = 1 - d[y][x-1]
			}
			if x != w-1 {
				d[y][x+1] = 1 - d[y][x+1]
			}
			cnt++
		}
		for i := 0; i < w; i++ {
			if f&(1<<i) != 0 {
				push(i, 0)
			}
		}
		for i := 1; i < h; i++ {
			for j := 0; j < w; j++ {
				if d[i-1][j] != 0 && v[i][j] != 0 {
					push(j, i)
				}
			}
		}
		ok = true
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if d[i][j] != 0 {
					ok = false
				}
			}
		}
		if ok {
			ans = min(ans, cnt)
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

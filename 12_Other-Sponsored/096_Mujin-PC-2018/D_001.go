package main

import "fmt"

var d [1000][1000]int

func f(a, b int) int {
	if d[a][b] != 0 {
		return d[a][b]
	}
	d[a][b] = -1
	c := min(a, b)
	if c > 99 {
		c += (c%10 - c/100) * 99
	} else if c > 9 {
		c += (c%10 - c/10) * 9
	}
	x := max(a, b)
	if c < x {
		x -= c
	} else {
		c -= x
	}
	if c == 0 || x == 0 {
		d[a][b] = 1
		return d[a][b]
	}
	d[a][b] = f(c, x)
	return d[a][b]
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	r := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if f(i, j) < 0 {
				r++
			}
		}
	}
	fmt.Println(r)
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

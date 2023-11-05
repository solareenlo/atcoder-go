package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var h, w, t int
	fmt.Scan(&h, &w, &t)

	stride := w + 1
	n := (h + 2) * stride

	var f func(int, int) int
	f = func(i, j int) int {
		return stride*(i+1) + (j + 1)
	}

	c := make([]int, n)
	for i := range c {
		c[i] = -1
	}
	var e [20]int
	k := 2
	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		for j := 0; j < w; j++ {
			if s[j] == '#' {
				c[f(i, j)] = 1
			} else {
				c[f(i, j)] = 0
			}
			if s[j] == 'S' {
				e[0] = f(i, j)
			}
			if s[j] == 'G' {
				e[1] = f(i, j)
			}
			if s[j] == 'o' {
				e[k] = f(i, j)
				k++
			}
		}
	}

	var x [20][20]int
	for u := 0; u < k; u++ {
		d := make([]int, n)
		for i := range d {
			d[i] = -1
		}
		q := make([]int, h*w)

		j0 := 0
		j1 := 0
		q[j1] = e[u]
		d[q[j1]] = 0
		j1++

		for j0 < j1 {
			i := q[j0]
			j0++
			d1 := d[i] + 1
			for _, l := range []int{-1, 1, -stride, stride} {
				i1 := i + l
				if c[i1] != 0 {
					continue
				}
				if d[i1] >= 0 {
					continue
				}
				d[i1] = d1
				q[j1] = i1
				j1++
			}
		}
		for i := 0; i < k; i++ {
			if d[e[i]] < 0 {
				x[u][i] = 1 << 29
			} else {
				x[u][i] = d[e[i]]
			}
		}
	}

	if x[0][1] < 0 || x[0][1] > t {
		fmt.Println(-1)
		return
	}

	var dp [1 << 18][20]int
	for j := 0; j < k; j++ {
		dp[0][j] = 1 << 29
	}
	dp[0][0] = 0
	for i := 1; i < 1<<(k-2); i++ {
		dp[i][0] = 1 << 29
		for h := 2; h < k; h++ {
			t := 1 << 29
			m := 1 << (h - 2)
			if (i & m) == 0 {
				dp[i][h] = t
				continue
			}
			i1 := i ^ m
			for j := 0; j < k; j++ {
				t = min(t, dp[i1][j]+x[h][j])
			}
			dp[i][h] = t
		}
		t := 1 << 29
		dp[i][1] = t
		for j := 0; j < k; j++ {
			t = min(t, dp[i][j]+x[1][j])
		}
		dp[i][1] = t
	}

	res := 0
	for i := 1; i < 1<<(k-2); i++ {
		if dp[i][1] <= t {
			res = max(res, popcount(uint32(i)))
		}
	}
	fmt.Println(res)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

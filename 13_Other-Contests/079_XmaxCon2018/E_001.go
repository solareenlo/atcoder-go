package main

import "fmt"

func main() {
	type pair struct {
		x, y int
	}

	var n int
	fmt.Scan(&n)
	if n == 1 {
		return
	}
	var lg [65545]int
	for i := 0; i <= n; i++ {
		lg[1<<i] = i
	}
	var suf [17]int
	cnt := 0
	for i := 1; i < n+1; i++ {
		cnt++
		suf[i] = cnt
	}
	var g [17][17][17]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := 0; k < n+1; k++ {
				cnt++
				g[i][j][k] = cnt
			}
		}
	}
	var s0, s1 [5]int
	for i := 0; i < 5; i++ {
		cnt++
		s0[i] = cnt
		cnt++
		s1[i] = cnt
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j >= 2; j-- {
			fmt.Printf("a[99999] = a[%d] AND in[%d]\n", suf[j-1], i)
			fmt.Printf("a[%d] = a[%d] OR a[99999]\n", suf[j], suf[j])
		}
		fmt.Printf("a[%d] = a[%d] OR in[%d]\n", suf[1], suf[1], i)
	}
	v := make([]pair, 0)
	num := 0
	v = append(v, pair{1, n})
	var rg [17][17]int
	for i := 1; i < n+1; i++ {
		rg[i][n] = suf[i]
	}
	for len(v) < n {
		nxt := make([]pair, 0)
		for i := 0; i < len(v); i++ {
			if v[i].x != v[i].y {
				l := v[i].x
				r := v[i].y
				mid := (l + r) / 2
				nxt = append(nxt, pair{l, mid})
				nxt = append(nxt, pair{mid + 1, r})
				fmt.Printf("a[%d] = a[%d] OR a[%d]\n", s1[num], s1[num], rg[mid+1][r])
			} else {
				nxt = append(nxt, v[i])
				fmt.Printf("a[%d] = a[%d] OR a[%d]\n", s1[num], s1[num], rg[v[i].x][v[i].x])
			}
		}
		fmt.Printf("a[%d] = NOT a[%d]\n", s0[num], s1[num])
		for i := 0; i < len(v); i++ {
			if v[i].x != v[i].y {
				l := v[i].x
				r := v[i].y
				mid := (l + r) / 2
				for i := l; i < mid+1; i++ {
					cnt++
					rg[i][mid] = cnt
					fmt.Printf("a[%d] = a[%d] AND a[%d]\n", rg[i][mid], rg[i][r], s0[num])
				}
			}
		}
		v = nxt
		num++
	}
	var f [17]int
	for i := 1; i < n+1; i++ {
		f[i] = rg[i][i]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			c := 0
			for k := 0; k < n; k++ {
				if k != i && k != j {
					c++
					for p := c; p >= 2; p-- {
						fmt.Printf("a[99999] = a[%d] AND in[%d]\n", g[i][j][p-1], k)
						fmt.Printf("a[%d] = a[%d] OR a[99999]\n", g[i][j][p], g[i][j][p])
					}
					fmt.Printf("a[%d] = a[%d] OR in[%d]\n", g[i][j][1], g[i][j][1], k)
				}
			}
			fmt.Printf("out[%d][%d] = a[%d] OR a[%d]\n", i, j, f[1], f[1])
			for k := 1; k < c+1; k++ {
				fmt.Printf("a[99999] = a[%d] AND a[%d]\n", g[i][j][k], f[k+1])
				fmt.Printf("out[%d][%d] = out[%d][%d] OR a[99999]\n", i, j, i, j)
			}
			fmt.Printf("a[99999] = in[%d] OR in[%d]\n", i, j)
			fmt.Printf("out[%d][%d] = out[%d][%d] AND a[99999]\n", i, j, i, j)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("out[%d][%d] = out[%d][%d] AND out[%d][%d]\n", i, j, j, i, j, i)
		}
	}
}

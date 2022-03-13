package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1505
const mod = 998244353

var (
	sum int
	c   = make([]int, N)
	vi  = make([]bool, N)
	e   = make([][]int, N)
)

func dfs(u int) {
	sum ^= c[u]
	vi[u] = true
	for _, v := range e[u] {
		if !vi[v] {
			dfs(v)
		}
	}
}

func solve(m int, a [N][N]int) int {
	res := 0
	for i := 1; i < m+1; i++ {
		for j := i + 1; j < m+1; j++ {
			if a[i][j] == 0 {
				res++
			}
		}
	}
	if m%2 == 0 {
		return res
	}
	for i := 1; i < m+1; i++ {
		e[i] = e[i][:0]
		vi[i] = false
		c[i] = 0
	}
	for i := 1; i < m+1; i++ {
		for j := i + 1; j < m+1; j++ {
			if a[i][j] == 1 {
				c[i] ^= 1
				c[j] ^= 1
			} else if a[i][j] == 0 {
				e[i] = append(e[i], j)
				e[j] = append(e[j], i)
			}
		}
	}
	for i := 1; i < m+1; i++ {
		if !vi[i] {
			sum = 0
			res++
			dfs(i)
			if sum != 0 {
				fmt.Println(0)
				os.Exit(0)
			}
		}
	}
	fmt.Println()
	return res - m
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	A := [2][N][N]int{}
	for i := 1; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		s = " " + s
		for j := 1; j < n; j++ {
			if s[j] == '?' {
				continue
			}
			k := (i + j) & 1
			a := i
			b := j
			t := i + j - n
			if a < b {
				a, b = b, a
			}
			if t > 0 {
				a -= t
				b -= t
			}
			aim := -1
			if s[j] == 'x' {
				aim = 1
			}
			u := (a-b)/2 + 1
			v := (a+b)/2 + 1
			if A[k][u][v] != 0 {
				if A[k][u][v] != aim {
					fmt.Println(0)
					os.Exit(0)
				}
				continue
			}
			A[k][u][v] = aim
		}
	}

	pw := [N * N]int{}
	pw[0] = 1
	for i := 1; i < n*n+1; i++ {
		pw[i] = pw[i-1] * 2 % mod
	}
	m0 := (n + 2) / 2
	m1 := (n + 1) / 2
	fmt.Println(pw[solve(m0, A[0])+solve(m1, A[1])])
}

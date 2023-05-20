package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var calc func(int, int, int) [][]int
	calc = func(n, m, p int) [][]int {
		rev := make([]int, p)
		for i := 1; i < p; i++ {
			for j := 1; j < p; j++ {
				v := i * j % p
				if v == 1 {
					rev[i] = j
				}
			}
		}
		coef := make([][]int, m)
		pcoef := make([][]int, m)
		for i := 0; i < m; i++ {
			coef[i] = make([]int, m)
			pcoef[i] = make([]int, m)
		}
		for i := 0; i < m; i++ {
			coef[i][i] = 1
		}
		for i := 0; i < n; i++ {
			ncoef := make([][]int, m)
			for j := range ncoef {
				ncoef[j] = make([]int, m)
			}
			for j := 0; j < m; j++ {
				for k := j - 1; k <= j+1; k++ {
					if k < 0 || k >= m {
						continue
					}
					z := 1
					if j != k {
						z = p - 1
					}
					for t := 0; t < m; t++ {
						ncoef[j][t] += z * coef[k][t]
						ncoef[j][t] %= p
					}
				}
				z := p - 1
				for t := 0; t < m; t++ {
					ncoef[j][t] += z * pcoef[j][t]
					ncoef[j][t] %= p
				}
			}
			pcoef, coef = coef, pcoef
			coef, ncoef = ncoef, coef
		}
		a := make([][]int, len(coef))
		for i := range a {
			a[i] = make([]int, len(coef[i]))
			copy(a[i], coef[i])
		}
		le := 0
		for j := 0; j < m; j++ {
			for i := le; i < m; i++ {
				if a[i][j] != 0 {
					a[i], a[le] = a[le], a[i]
					break
				}
			}
			if le < m && a[le][j] != 0 {
				c := rev[a[le][j]]
				for k := 0; k < m; k++ {
					a[le][k] = a[le][k] * c % p
				}
				for i := le + 1; i < m; i++ {
					c := (p - a[i][j]) % p
					for k := 0; k < m; k++ {
						a[i][k] += c * a[le][k]
						a[i][k] %= p
					}
				}
				le++
			}
		}
		res := make([][]int, 0)
		if le == m {
			return res
		}
		res = make([][]int, n)
		for i := range res {
			res[i] = make([]int, m)
		}
		val := make([]int, m)
		for i := range val {
			val[i] = 1
		}
		for i := le - 1; i >= 0; i-- {
			id := 0
			for j := 0; j < m; j++ {
				if a[i][j] != 0 {
					id = j
					break
				}
			}
			v := 0
			for j := id + 1; j < m; j++ {
				v += a[i][j] * val[j]
			}
			v %= p
			v = (p - v) % p
			val[id] = v
		}
		for j := 0; j < m; j++ {
			res[0][j] = val[j]
		}
		for i := 0; i < n-1; i++ {
			for j := 0; j < m; j++ {
				v := res[i][j]
				if j > 0 {
					v += p - res[i][j-1]
				}
				if j+1 < m {
					v += p - res[i][j+1]
				}
				if i > 0 {
					v += p - res[i-1][j]
				}
				res[i+1][j] = v % p
			}
		}
		return res
	}

	var solve func()
	solve = func() {
		var n, m, x int
		fmt.Fscan(in, &n, &m, &x)
		cop := x
		for i := 2; i <= cop; i++ {
			if cop%i == 0 {
				p := i
				for cop%p == 0 {
					cop /= p
				}
				ans := calc(n, m, p)
				if len(ans) != 0 {
					for i := 0; i < n; i++ {
						for j := 0; j < m; j++ {
							ans[i][j] *= x / p
						}
					}
					for i := 0; i < n; i++ {
						for j := 0; j < len(ans[i]); j++ {
							if j > 0 {
								fmt.Fprint(out, " ")
							}
							fmt.Fprint(out, ans[i][j])
						}
						fmt.Fprintln(out)
					}
					return
				}
			}
		}
		fmt.Fprintln(out, -1)
	}

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solve()
	}
}

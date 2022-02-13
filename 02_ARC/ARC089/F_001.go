package main

import "fmt"

const N = 75
const mod = 1_000_000_007

var (
	n   int
	k   int
	m   int
	s   string
	ans int = 1
	x       = [N]int{}
	a       = [N]int{}
	C       = [2 * N][2 * N]int{}
)

func dfs(d, l, r, chose int) {
	if chose != 0 {
		R := 1
		B := 1
		S := 0
		for i := 1; i <= k; i++ {
			if s[i] == 'r' {
				if R <= m {
					R++
				} else if S != 0 {
					S--
				}
			}
			if s[i] == 'b' {
				if B < R && x[B] != 0 {
					S += x[B] - 1
					B++
				} else if S != 0 {
					S--
				}
			}
		}
		if R > m && (B > m || x[B] == 0) && S == 0 {
			ret := C[n-l+r][r]
			for i := d + 1; i <= x[1]; i++ {
				S += a[i]
				ret = ret * C[S][a[i]] % mod
			}
			ans = (ans + ret) % mod
		}
	}
	if d < 0 {
		return
	}
	tm := m
	tl := 1
	if d != 0 {
		tl = 2*d - 1
	}
	tr := 2*d + 1
	for i := 0; l <= n; i++ {
		a[d] = i
		dfs(d-1, l, r, i)
		m++
		x[m] = d
		tmp := 0
		if l != 0 {
			tmp = 1
		}
		l += tl + tmp
		r += tr + 1
	}
	m = tm
}

func main() {
	fmt.Scan(&n, &k, &s)
	s = " " + s

	C[0][0] = 1
	for i := 1; i < 150; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % mod
		}
	}

	dfs(n, 0, 0, 0)

	fmt.Println(ans)
}

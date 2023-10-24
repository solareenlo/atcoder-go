package main

import "fmt"

const N = 32
const M = 100

var n, m int
var T int
var mp map[int]int
var rev [M]int
var to [N]int
var mat [M][M]int

func dfs(u int) int {
	if mp[u] != 0 {
		return mp[u]
	}
	T++
	mp[u] = T
	x := mp[u]
	rev[T] = u
	for H := 0; H < 1<<n; H++ {
		v := 0
		for S := 0; S < 1<<n; S++ {
			if ((u>>S)&1) != 0 && (S&H) == S {
				v |= to[H^S]
			}
		}
		if v != 0 {
			mat[x][dfs(v)]++
		}
	}
	return x
}

func main() {
	const P = 998244353

	var f, g [M]int
	mp = make(map[int]int)

	fmt.Scan(&n, &m)
	for S := 0; S < 1<<n; S++ {
		to[S] = 1 << S
		for i := 0; i+1 < n; i++ {
			if ((S>>i)&1) != 0 && ((S>>(i+1))&1) != 0 {
				to[S] |= to[S^(1<<i)^(1<<(i+1))]
			}
		}
	}
	dfs(1)
	f[1] = 1
	for m > 0 {
		if (m & 1) != 0 {
			for i := range g {
				g[i] = 0
			}
			f, g = g, f
			for i := 1; i <= T; i++ {
				for j := 1; j <= T; j++ {
					f[i] = (f[i] + g[j]*mat[j][i]%P) % P
				}
			}
		}
		var tmp [M][M]int
		mat, tmp = tmp, mat
		for i := 1; i <= T; i++ {
			for j := 1; j <= T; j++ {
				for k := 1; k <= T; k++ {
					mat[i][j] = (mat[i][j] + tmp[i][k]*tmp[k][j]%P) % P
				}
			}
		}
		m >>= 1
	}
	ans := 0
	for i := 1; i <= T; i++ {
		if (rev[i] & 1) != 0 {
			ans = (ans + f[i]) % P
		}
	}
	fmt.Println(ans)
}

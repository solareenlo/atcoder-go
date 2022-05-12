package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353
const N = 100000

var (
	pri = [100100]int{}
	phi = [100100]int{}
)

func sieve() {
	phi[1] = 1
	for i := 2; i <= N; i++ {
		if pri[i] == 0 {
			pri[0]++
			pri[pri[0]] = i
			phi[i] = i - 1
		}
		for j := 1; j <= pri[0] && i*pri[j] <= N; j++ {
			pri[i*pri[j]] = 1
			if !(i%pri[j] != 0) {
				phi[i*pri[j]] = phi[i] * pri[j]
				break
			}
			phi[i*pri[j]] = phi[i] * phi[pri[j]]
		}
	}
}

var (
	sp  = [100100]bool{}
	tot int
	f   = [100100]int{}
	g   = [100100]int{}
	v   = make([][]int, 100100)
)

func dfs(x, fa int) {
	sp[x] = false
	f[x] = 1
	g[x] = 1
	for _, y := range v[x] {
		if y != fa && sp[y] {
			dfs(y, x)
			tot = (f[x]*g[y] + f[y]*g[x] + tot) % mod
			f[x] += f[y]
			f[x] %= mod
			g[x] = (f[y] + g[x] + g[y]) % mod
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	sieve()

	a := make([]int, n+1)
	u := make([][]int, 100100)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		u[a[i]] = append(u[a[i]], i)
	}

	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}

	res := 0
	for i := 1; i <= N; i++ {
		for j := i; j <= N; j += i {
			for _, k := range u[j] {
				sp[k] = true
			}
		}
		for j := i; j <= N; j += i {
			for _, k := range u[j] {
				if sp[k] {
					dfs(k, 0)
				}
			}
		}
		res = (phi[i]*tot + res) % mod
		tot = 0
	}
	fmt.Println(res)
}

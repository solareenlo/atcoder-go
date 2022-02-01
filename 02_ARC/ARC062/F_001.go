package main

import "fmt"

const N = 53
const M = 103

var k int

func calc(s int) int {
	t := 0
	for i := 1; i <= s; i++ {
		t += powMod(k, gcd(s, i))
		t %= mod
	}
	return t * powMod(s, mod-2) % mod
}

type edge struct{ to, nxt int }

var (
	e   = make([]edge, M<<1)
	lst = make([]int, N)
	tt  int
)

func link(u, v int) {
	tt++
	e[tt].to = v
	e[tt].nxt = lst[u]
	lst[u] = tt
}

var (
	dfn = [N]int{}
	low = [N]int{}
	ti  int
	vv  int
	ee  int
	ans int = 1
)

func tarjan(u int) {
	ti++
	dfn[u] = ti
	low[u] = ti
	vv++
	for i := lst[u]; i > 0; i = e[i].nxt {
		v := e[i].to
		if dfn[v] == 0 {
			ne := ee
			nv := vv
			tarjan(v)
			low[u] = min(low[u], low[v])
			if low[v] >= dfn[u] {
				t := ee - ne
				if vv-nv+1 == t {
					ans = ans * calc(t) % mod
				} else {
					ans = ans * nCrMod(t+k-1, k-1) % mod
				}
				ee = ne
				vv = nv
			}
		} else {
			if dfn[u] > dfn[v] {
				ee++
			}
			low[u] = min(low[u], dfn[v])
		}
	}
}

func main() {
	initMod()
	var n, m int
	fmt.Scan(&n, &m, &k)

	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		link(x, y)
		link(y, x)
	}
	for i := 1; i <= n; i++ {
		if dfn[i] == 0 {
			tarjan(i)
		}
	}
	ans = ans * powMod(k, ee)
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

const mod = 1000000007
const size = 101010

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, mod-2)
}

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}

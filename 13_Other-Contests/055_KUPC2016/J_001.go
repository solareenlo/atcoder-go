package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

type pair struct {
	x, y int
}

const N = 100010

var K int
var rd, ifac, par, in [N]int
var sz, hs, way [N]int
var ne [N]int
var G [N][]int

func main() {
	IN := bufio.NewReader(os.Stdin)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		rd[i] = rand.Int() % P
	}
	for i := 0; i < N; i++ {
		ifac[i] = 1
		if i >= 2 {
			ifac[i] = P - (P/i)*ifac[P%i]%P
		}
	}
	for i := 1; i < N; i++ {
		ifac[i] = ifac[i-1] * ifac[i] % P
	}
	var n int
	fmt.Fscan(IN, &n, &K)
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &par[i])
		par[i]--
		G[par[i]] = append(G[par[i]], i)
		in[par[i]]++
	}
	pts := make([]int, 0)
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			pts = append(pts, i)
		}
	}
	for i := 0; i < len(pts); i++ {
		in[par[pts[i]]]--
		if in[par[pts[i]]] == 0 {
			pts = append(pts, par[pts[i]])
		}
	}
	cir := make([]int, 0)
	s := 0
	for in[s] == 0 {
		s++
	}
	for t := par[s]; ; {
		cir = append(cir, t)
		if t == s {
			break
		}
		t = par[t]
	}
	for _, c := range cir {
		for i := 0; i < len(G[c]); i++ {
			if in[G[c][i]] != 0 {
				G[c] = erase(G[c], i)
				break
			}
		}
		Dfs(c)
	}
	up := 0
	down := 0
	L := len(cir)
	smallloop := kmp(cir)
	block := L / smallloop
	preprod := make([]int, L)
	for i := 0; i < L; i++ {
		preprod[i] = way[cir[i]] % P
		if i != 0 {
			preprod[i] = preprod[i-1] * way[cir[i]] % P
		}
	}
	for step := 1; step < block+1; step++ {
		orbit := gcd(step, block)
		up = (up + preprod[smallloop*orbit-1]) % P
		down++
	}
	v := make([]pair, 0)
	times := make(map[pair]int)
	for i := 0; i < L; i++ {
		v = append(v, pair{hs[cir[i]], way[cir[i]]})
		v = append(v, pair{-1, 1})
	}
	for _, e := range v {
		times[e]++
	}
	cof := 1
	for k, v := range times {
		t := v / 2
		for i := 0; i < t; i++ {
			cof = cof * k.y % P
		}
	}
	v = append(v, v...)
	length := manacher(v)
	for i := 0; i < L; i++ {
		if length[i+L] >= L+1 {
			x := v[i+L]
			y := v[i+L+L]
			newcof := cof
			times[x]++
			if (^times[x] & 1) != 0 {
				newcof = newcof * x.y % P
			}
			times[y]++
			if (^times[y] & 1) != 0 {
				newcof = newcof * y.y % P
			}
			times[x]--
			times[y]--
			up = (up + newcof) % P
			down++
		}
	}
	fmt.Println(up * powMod(down, P-2) % P)
}

const P = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % P
		}
		a = a * a % P
		n /= 2
	}
	return res
}

func erase(a []int, pos int) []int {
	return append(a[:pos], a[pos+1:]...)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Dfs(c int) {
	sz[c] = 1
	hs[c] = 1
	way[c] = K
	for _, t := range G[c] {
		Dfs(t)
		sz[c] += sz[t]
		hs[c] = hs[c] * (rd[sz[t]] + hs[t]) % P
	}
	sort.Slice(G[c], func(a, b int) bool {
		return hs[G[c][a]] < hs[G[c][b]]
	})
	for i, j := 0, 0; i < len(G[c]); i = j {
		for j = i; j < len(G[c]) && hs[G[c][i]] == hs[G[c][j]]; j++ {

		}
		way[c] = way[c] * C(j-i+way[G[c][i]]-1, j-i) % P
	}
}

func kmp(v []int) int {
	ne[0] = -1
	n := len(v)
	for i, j := 1, -1; i < n; i++ {
		for j != -1 && hs[v[j+1]] != hs[v[i]] {
			j = ne[j]
		}
		ne[i] = j
		if hs[v[j+1]] == hs[v[i]] {
			j++
			ne[i] = j
		}
	}
	if n%(n-1-ne[n-1]) == 0 {
		return n - 1 - ne[n-1]
	}
	return n
}

func C(a, b int) int {
	r := 1
	for i := 0; i < b; i++ {
		r = r * (a - i) % P
	}
	return r * ifac[b] % P
}

func manacher(s []pair) []int {
	n := len(s)
	p := make([]int, n)
	for i := range p {
		p[i] = 1
	}
	for i, j, k := 1, 0, 1; i < n; i++ {
		if i < k {
			p[i] = min(p[2*j-i], k-i)
		}
		for i-p[i] >= 0 && i+p[i] < n && s[i-p[i]] == s[i+p[i]] {
			p[i]++
		}
		if k < i+p[i] {
			k = i + p[i]
			j = i
		}
	}
	return p
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

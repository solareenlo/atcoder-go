package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const N = 21
const M = 1 << 20

type pair struct{ x, y int }

var (
	e = make([][]pair, N)
	w = make([]int, N)
)

func dfs(u, fa int) {
	for _, v := range e[u] {
		if v.x == fa {
			continue
		}
		w[v.x] = w[u] | (1 << v.y)
		dfs(v.x, u)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	initMod()
	n--

	f := make([]int, M)
	for i := 1; i < n+1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		e[a] = append(e[a], pair{b, i - 1})
		e[b] = append(e[b], pair{a, i - 1})
		f[1<<(i-1)]++
	}

	dfs(1, 0)

	for i := n + 1; i < m+1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		f[w[a]^w[b]]++
	}

	for k, d := 0, 1; k < n; k++ {
		for i := 0; i < (1 << n); i += d << 1 {
			l := f[i:]
			r := f[i+d:]
			for j := 0; j < d; j++ {
				r[j] += l[j]
			}
		}
		d <<= 1
	}

	cnt := make([]int, M)
	cnt[(1<<n)-1] = 1
	sum := make([]int, M)

	for s := (1 << n) - 1; s > 0; s-- {
		for i := 0; i < n; i++ {
			if (s >> i & 1) == 0 {
				continue
			}
			t := s ^ (1 << i)
			C := cnt[s]
			S := sum[s]
			num := m - f[s]
			nw := f[s] - f[t] - 1
			C = mul(C, mul(fact[num+nw], invf[num]))
			S = mul(S, mul(fact[num+nw+1], invf[num+1]))
			S = add(S, mul(C, n-bits.OnesCount(uint(t))))
			cnt[t] = add(cnt[t], C)
			sum[t] = add(sum[t], S)
		}
	}
	fmt.Println(sum[0])
}

func mul(x, y int) int {
	return x * y % mod
}

func add(x, y int) int {
	if x+y >= mod {
		return x + y - mod
	}
	return x + y
}

const mod = 1000000007
const size = 1 << 20

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

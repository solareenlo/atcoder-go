package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353

func hadamard(a []int) []int {
	n := len(a)
	for i := 1; i < n; i <<= 1 {
		for j := 0; j < n; j++ {
			if j&i != 0 {
				k := j ^ i
				tmp := a[j]
				a[j] = (a[k] - a[j] + mod) % mod
				a[k] += tmp
				a[k] %= mod
			}
		}
	}
	return a
}

type pair struct{ x, y int }

func dfs(x, t int) pair {
	if t == 0 {
		return pair{0, 1}
	}
	p := dfs(x, t>>1)
	p.x += p.x * p.y % mod
	p.x %= mod
	p.y *= p.y % mod
	p.y %= mod
	if t&1 != 0 {
		p.y *= x
		p.y %= mod
		p.x += p.y
		p.x %= mod
	}
	return p
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	m2 := 1 << 16
	d := make([]int, m2)
	for i := 0; i < k; i++ {
		var a int
		fmt.Fscan(in, &a)
		d[a] += 1
	}

	d = hadamard(d)
	for i := 0; i < m2; i++ {
		d[i] = dfs(d[i], n).x
	}
	d = hadamard(d)

	for i := 0; i < m2; i++ {
		d[i] = divMod(d[i], m2)
	}

	res := 0
	for i := 0; i < m2; i++ {
		if i != 0 {
			res += d[i]
			res %= mod
		}
	}
	fmt.Println(res)
}

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

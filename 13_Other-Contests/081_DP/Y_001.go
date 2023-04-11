package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 3010

	type pair struct {
		x, y int
	}

	var h, w, n int
	fmt.Fscan(in, &h, &w, &n)

	p := make([]pair, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	tmp := p[:n]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
	p[n].x = h
	p[n].y = w
	initMod()
	var f [N]int
	for i := 0; i <= n; i++ {
		f[i] = nCrMod(p[i].x+p[i].y-2, p[i].x-1)
		for j := 0; j < i; j++ {
			if p[j].y <= p[i].y {
				f[i] = (f[i] - f[j]*nCrMod(p[i].x-p[j].x+p[i].y-p[j].y, p[i].x-p[j].x)%mod + mod) % mod
			}
		}
	}
	fmt.Println(f[n])
}

const mod = 1000000007
const size = 200010

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

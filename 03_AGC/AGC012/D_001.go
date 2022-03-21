package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)

	d := make([]int, size)
	for i := 0; i < size; i++ {
		d[i] = 1 << 60
	}
	c := make([]int, size)
	w := make([]int, size)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i], &w[i])
		d[c[i]] = min(d[c[i]], w[i])
	}

	mn := 1 << 60
	mn2 := 1 << 60
	for i := 1; i <= n; i++ {
		if d[i] < mn {
			mn2 = mn
			mn = d[i]
		} else if d[i] < mn2 {
			mn2 = d[i]
		}
	}

	if mn+mn2 > y {
		fmt.Println(1)
		return
	}

	initMod()

	f := make([]int, size)
	for i := 1; i <= n; i++ {
		tmp := mn
		if d[c[i]] == mn {
			tmp = mn2
		}
		if (w[i]+d[c[i]] <= x) || (w[i]+tmp <= y) {
			f[c[i]]++
		}
	}

	t := 0
	ans := 1
	for i := 1; i <= n; i++ {
		if d[i]+mn <= y {
			t += f[i]
			ans = ans * invf[f[i]] % mod
		}
	}

	fmt.Println(ans * fact[t] % mod)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mod = 1000000007
const size = 200007

var fact, invf, fciv [size]int

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

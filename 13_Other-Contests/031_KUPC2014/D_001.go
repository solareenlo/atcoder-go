package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	initMod()

	var s1, s2 string
	var d1, d2 int
	fmt.Fscan(in, &s1, &d1, &s2, &d2)
	same, diff := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			same++
		} else {
			diff++
		}
	}
	k := d1 + d2 - diff
	if k > same*2 || k%2 != 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(nCrMod(same, k/2) * nCrMod(diff, d1-k/2) % mod)
}

const mod = 1000000007
const size = 200000

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

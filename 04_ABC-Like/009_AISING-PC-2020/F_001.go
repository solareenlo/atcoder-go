package main

import (
	"bufio"
	"fmt"
	"os"
)

func C(n, k int) int {
	if n < k {
		return 0
	}
	if k == 0 {
		return 1
	}
	r := 1
	for i := 0; i < k; i++ {
		r = r * (n - i) % mod * invMod(i+1) % mod
	}
	return r
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		res := 0
		for i := (n + 1) % 2; i <= 11; i += 2 {
			res += C(11, i) * C((n-i+25)/2, 15) % mod
		}
		fmt.Println(res % mod)
	}
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
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

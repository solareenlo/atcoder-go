package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)

	a := make([]int, n)
	c := make([]int, q+2)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i < q+1; i++ {
		fmt.Fscan(in, &c[i])
		c[i]--
	}

	d := make([]int, n)
	for i := 1; i < n; i++ {
		d[i] = (d[i-1] + powMod(a[i-1], a[i])) % mod
	}

	ans := 0
	for i := 0; i < q+1; i++ {
		if c[i] < c[i+1] {
			ans = (ans + (d[c[i+1]]-d[c[i]]+mod)%mod) % mod
		} else {
			ans = (ans + (d[c[i]]-d[c[i+1]]+mod)%mod) % mod
		}
	}
	fmt.Println(ans)
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

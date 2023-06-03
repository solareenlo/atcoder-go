package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, p int
	fmt.Fscan(in, &n, &p)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[i] = x % mod
	}

	ans := 0
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			if p == 0 {
				ans += i
			}
		} else {
			x := divMod(p, a[i])
			ans = ans + cnt[x]
		}
		cnt[a[i]]++
	}
	fmt.Println(ans)
}

const mod = 1000000007

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

package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for _, nx := range s {
			v[i] |= (1 << (nx - 'a'))
		}
	}

	res := 0
	for i := 1; i < (1 << n); i++ {
		ch := (1 << 26) - 1
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				ch &= v[j]
			}
		}
		pc := bits.OnesCount(uint(ch))
		if bits.OnesCount(uint(i))%2 != 0 {
			res += powMod(pc, k)
			res %= mod
		} else {
			res += (mod - powMod(pc, k))
			res %= mod
		}
	}

	fmt.Println(res)
}

const mod = 998244353

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

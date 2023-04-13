package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	initMod()

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var n, k int
		fmt.Fscan(in, &n, &k)
		ans := 0
		k = n*n - 8*k
		var t int
		if k >= 0 {
			t = int(math.Floor(math.Sqrt(float64(k))))
		}
		if t*t == k {
			if (n+t)%4 == 0 {
				ans = nCrMod(n/2, (n+t)/4)
			} else if (n-t)%4 == 0 {
				ans = nCrMod(n/2, (n-t)/4)
			}
			if k != 0 {
				ans = ans * 2 % mod
			}
		}
		fmt.Fprintln(out, ans)
	}
}

const mod = 1000000007
const size = 100005

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

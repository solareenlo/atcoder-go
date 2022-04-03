package main

import (
	"fmt"
	"math"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	high := make([]int, 50000)
	low := make([]int, 50000)
	L := int(math.Round(math.Sqrt(float64(M))))
	K := M / (L + 1)
	for i := 1; i <= L+K; i++ {
		x := M / (L + K + 1 - i)
		if i <= L {
			x = i
		}
		sum := 0
		for l, r, le := 2, 0, min(N, x); l <= le; l = r {
			q := x / l
			r = min(N, x/q) + 1
			var cur int
			if q <= L {
				cur = low[q]
			} else {
				cur = high[M/q]
			}
			sum += cur * (r - l) % mod
			sum %= mod
		}
		if i <= L {
			low[i] = divMod(sum+N%mod, N-1)
		} else {
			high[L+K+1-i] = divMod(sum+N%mod, N-1)
		}
	}

	if M <= L {
		fmt.Println(low[M])
	} else {
		fmt.Println(high[1])
	}
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

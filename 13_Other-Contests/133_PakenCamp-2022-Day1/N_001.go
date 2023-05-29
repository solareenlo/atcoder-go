package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf = 1_000_000_000_000_000

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b [3]int

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var s, t, P int
		fmt.Fscan(in, &s, &t, &P)
		for i := 0; i < 3; i++ {
			fmt.Fscan(in, &a[i], &b[i])
		}
		ans := inf
		for i := 0; i < 3; i++ {
			C := a[0] * a[1] % P * a[2] % P
			D := (a[1]*a[2]%P*b[0]%P + a[2]*b[1]%P + b[2]) % P
			val := solve(s, t, C, D, P)
			if val != -1 {
				ans = min(ans, 3*val+i)
			}
			s = (a[0]*s + b[0]) % P
			x := a[0]
			y := b[0]
			a[0] = a[1]
			b[0] = b[1]
			a[1] = a[2]
			b[1] = b[2]
			a[2] = x
			b[2] = y
		}
		if ans == inf {
			fmt.Println(-1)
		} else {
			fmt.Println(ans)
		}
	}
}

func solve(s, t, C, D, P int) int {
	if C == 0 {
		if t == s {
			return 0
		} else if t == D {
			return 1
		} else {
			return -1
		}
	}
	if C == 1 {
		if D == 0 {
			if t == s {
				return 0
			} else {
				return -1
			}
		} else {
			return (t - s + P) * invMod(D, P) % P
		}
	}
	inve := invMod((1 - C + P), P)
	inve = inve * D % P
	t = (t + (P - inve)) % P
	s = (s + (P - inve)) % P
	if s == 0 {
		if t == 0 {
			return 0
		} else {
			return -1
		}
	} else {
		t = t * invMod(s, P) % P
		return BSGS(C, t, P)
	}
}

func BSGS(a, b, prime int) int {
	mp := make(map[int]int)
	var i int
	val := 1
	for i = 0; i*i < prime; i++ {
		if _, ok := mp[val]; !ok {
			mp[val] = i
		}
		val = val * a % prime
	}
	lim := i
	inv := 1
	ex := prime - 2
	for ex != 0 {
		if (ex & 1) != 0 {
			inv = inv * val % prime
		}
		val = val * val % prime
		ex /= 2
	}
	for j := 0; j*lim < prime; j++ {
		if _, ok := mp[b]; ok {
			return j*lim + mp[b]
		}
		b = b * inv % prime
	}
	return -1
}

func powMod(a, n, mod int) int {
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

func invMod(a, mod int) int {
	return powMod(a, mod-2, mod)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

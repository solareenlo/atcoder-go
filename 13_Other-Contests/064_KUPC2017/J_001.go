package main

import (
	"bufio"
	"fmt"
	"os"
)

const MX = 10005

var A [55][MX]int
var B [MX][55]int
var C [55][55]int
var D, E [MX]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)
	p := 0
	var uf uf
	for i := 1; i <= H; i++ {
		var buf string
		fmt.Fscan(in, &buf)
		buf = " " + buf
		for j := 1; j <= W; j++ {
			if buf[j] != '#' {
				C[j][j]++
				D[i]++
				A[j][i] = 1
				B[i][j] = mod - 1
				p += uf.merge(j, i+W)
			}
		}
	}
	if p != H+W-1 {
		fmt.Println(-1)
		return
	}

	for i := 1; i <= H; i++ {
		E[i] = powMod(D[i], mod-2)
	}
	for i := H - 1; i >= 1; i-- {
		for j := 1; j <= W; j++ {
			if A[j][i] != 0 {
				for k := 1; k <= W; k++ {
					C[j][k] = (C[j][k] + E[i]*B[i][k]) % mod
				}
				A[j][i] = 0
			}
		}
	}

	ch := 1
	for i := W; i >= 1; i-- {
		if C[i][i] == 0 {
			ch *= -1
			for j := i - 1; j >= 1; j-- {
				if C[j][i] != 0 {
					for k := 1; k <= W; k++ {
						C[j][k], C[i][k] = C[i][k], C[j][k]
					}
				}
			}
		}
		piv := powMod(C[i][i], mod-2)
		for j := i - 1; j >= 1; j-- {
			mul := piv * (mod - C[j][i]) % mod
			for k := 1; k <= W; k++ {
				C[j][k] = (C[j][k] + mul*C[i][k]) % mod
			}
		}
	}

	ans := (ch + mod) % mod
	for i := 1; i <= W; i++ {
		ans = ans * C[i][i] % mod
	}
	for i := 1; i <= H-1; i++ {
		ans = ans * D[i] % mod
	}
	fmt.Println(H+W-1, ans)
}

type uf struct {
	t [MX + 55]int
}

func (u *uf) find(x int) int {
	if u.t[x] != 0 {
		u.t[x] = u.find(u.t[x])
		return u.t[x]
	}
	return x
}

func (u *uf) merge(a, b int) int {
	a = u.find(a)
	b = u.find(b)
	if a == b {
		return 0
	}
	u.t[a] = b
	return 1
}

const mod = 1000000007

func powMod(a, n int) int {
	a = a%mod + mod
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

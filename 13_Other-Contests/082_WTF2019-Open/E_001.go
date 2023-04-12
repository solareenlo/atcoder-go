package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

var P [2][1111]int
var r int
var q [3]int
var c [2]int
var f [1111]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	r = 1
	if s[0] != 'X' {
		P[0][0] = 1
		c[0] = 1
		r = mod - 1
	}
	d := 1
	f[0] = 1
	for i := 1; i < n; i++ {
		f[i] = f[i-1] * inv(i, mod) % mod
		P[0][d] = r
		P[1][d] = 0
		d++
		r = (mod - r) % mod
		if s[i] != 'X' {
			for j := 0; j < 2; j++ {
				c[j] = (P[j][d-1] + r*j) % mod
				for k := 1; k < d; k++ {
					c[j] = (c[j] + P[j][d-1-k]*f[k]) % mod
				}
			}
			if s[i-1] != 'X' {
				d = 1
				P[0][0] = 0
				P[1][0] = 0
				r = 0
			}
			continue
		}
		for j := 0; j < 2; j++ {
			P[j][d-1] = (P[j][d-1] + c[j]) % mod
			P[j][d-2] = (P[j][d-2] + mod - c[j]) % mod
		}
		if s[i-1] == 'X' {
			fmt.Println(0, 0, 0)
			return
		}
	}
	q[0] = r * inv(2, mod) % mod
	q[2] = (mod - q[0]) % mod
	for i := 0; i < d; i++ {
		for j := 0; j < 2; j++ {
			q[j] = (q[j] + P[j][d-1-i]) % mod
			q[j+1] = (q[j+1] + P[j][d-1-i]*(mod-f[i])) % mod
		}
		f[i+1] = (f[i+1] + f[i]) % mod
	}
	if s[n-1] != 'X' {
		for j := 0; j < 2; j++ {
			q[j+1] = (q[j+1] + c[j]) % mod
		}
	}
	fmt.Println(q[0], q[1], q[2])
}

func inv(a, m int) int {
	i, j := 1, 0
	for b := m; a > 1; {
		j = (j + b/a*(m-i)) % m
		b %= a
		i, j = j, i
		a, b = b, a
	}
	return i
}

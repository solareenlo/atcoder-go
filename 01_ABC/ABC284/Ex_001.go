package main

import "fmt"

var k, m, n, s int
var p2 [1010]int
var a, b, p1, p3, p4 [32]int
var q1, q2 [32][32]int

func f(x, y int) {
	if y == 0 {
		if x < k {
			return
		}
		t := 0
		for i := 1; i <= x; i++ {
			t += a[i] / 2
		}
		for i := 1; i < x; i++ {
			for j := i + 1; j <= x; j++ {
				t += q2[a[i]][a[j]]
			}
		}
		t = p2[t] * q1[x][k] % m * p3[k] % m
		for i := 1; i <= x; i++ {
			t = t * p1[a[i]] % m
		}
		for i := 1; i <= n; i++ {
			t = t * p4[b[i]] % m
		}
		s += t
	} else {
		for i := a[x]; i <= y; i++ {
			a[x+1] = i
			b[i]++
			f(x+1, y-i)
			b[i]--
		}
	}
}

func main() {
	fmt.Scan(&n, &k, &m)
	p1[1] = 1
	p2[0] = 1
	p3[0] = 1
	p4[0] = 1
	q1[1][1] = 1
	a[0] = 1
	for i := 2; i <= n; i++ {
		p1[i] = (m - m/i) * p1[m%i] % m
	}
	for i := 1; i <= 1000; i++ {
		p2[i] = p2[i-1] * 2 % m
	}
	for i := 1; i <= n; i++ {
		p3[i] = p3[i-1] * i % m
		p4[i] = p4[i-1] * p1[i] % m
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			q1[i][j] = (q1[i-1][j-1] + j*q1[i-1][j]) % m
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			q2[i][j] = gcd(i, j)
		}
	}
	f(0, n)
	fmt.Println(s % m)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

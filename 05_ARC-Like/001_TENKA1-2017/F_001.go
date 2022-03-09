package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func exgcd(a, b int, x, y *int) {
	if b == 0 {
		*x = 1
		*y = 0
		return
	}
	exgcd(b, a%b, y, x)
	*y -= a / b * *x
}

func phi(n int) int {
	s := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			s = s / i * (i - 1)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		s = s / n * (n - 1)
	}
	return s
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

func solve(a, m int) int {
	if m == 1 {
		return 1
	}
	p := phi(m)
	g := gcd(m, p)
	t := solve(a, g)
	ay := powMod(a, t, m)
	var x, y int
	exgcd(p, m, &x, &y)
	ay -= t % p
	if ay < 0 {
		x = -x
		ay = -ay
	}
	x = ((ay/g*x%m+m)%m)*p + t%p
	t = m / g * p
	if x < t {
		x += t
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		var a, m int
		fmt.Fscan(in, &a, &m)
		fmt.Fprintln(out, solve(a, m))
	}
}

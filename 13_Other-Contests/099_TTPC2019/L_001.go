package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var M int

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var K int
	var s string
	fmt.Fscan(in, &K, &s)

	M = 1
	for i := 0; i < K; i++ {
		M *= 10
	}

	ans := 0
	d := div_ceil(K, 2)
	lim := 1
	for i := 0; i < d; i++ {
		lim *= 10
	}
	for q := 0; q < lim; q++ {
		i := 0
		tmp := expr(s, &i, q)
		b := tmp.x
		a := tmp.y
		a = a * lim % M
		b = (M - b) % M
		g := gcd(a, M)
		if b%g == 0 {
			ans += g / lim
		}
	}
	fmt.Println(ans)
}

func expr(s string, i *int, x int) pair {
	ret := term(s, i, x)
	for *i < len(s) && (s[*i] == '+' || s[*i] == '-') {
		op := s[*i]
		(*i)++
		rhs := term(s, i, x)
		if op == '+' {
			ret = add(ret, rhs)
		} else {
			ret = sub(ret, rhs)
		}
	}
	return ret
}

func term(s string, i *int, x int) pair {
	ret := factor(s, i, x)
	for *i < len(s) && s[*i] == '*' {
		(*i)++
		rhs := factor(s, i, x)
		ret = mul(ret, rhs)
	}
	return ret
}

func factor(s string, i *int, x int) pair {
	ret := value(s, i, x)
	if *i < len(s) && s[*i] == '^' {
		(*i)++
		p := number(s, i, x)
		ret = pow(ret, p.x)
	}
	return ret
}

func value(s string, i *int, x int) pair {
	if unicode.IsDigit(rune(s[*i])) {
		return number(s, i, x)
	} else if s[*i] == 'x' {
		(*i)++
		return pair{x, 1}
	}
	(*i)++
	ret := expr(s, i, x)
	(*i)++
	return ret
}

func number(s string, i *int, x int) pair {
	ret := 0
	for *i < len(s) && unicode.IsDigit(rune(s[*i])) {
		ret = ret*10 + int(s[*i]-'0')
		(*i)++
	}
	return pair{ret, 0}
}

func add(lhs, rhs pair) pair {
	return pair{(lhs.x + rhs.x) % M, (lhs.y + rhs.y) % M}
}

func sub(lhs, rhs pair) pair {
	return pair{(lhs.x - rhs.x + M) % M, (lhs.y - rhs.y + M) % M}
}

func mul(lhs, rhs pair) pair {
	return pair{(lhs.x * rhs.x) % M, (lhs.x*rhs.y + lhs.y*rhs.x) % M}
}

func pow(lhs pair, n int) pair {
	return pair{PowMod(lhs.x, n), (n * PowMod(lhs.x, n-1) * lhs.y) % M}
}

func div_ceil(a, b int) int {
	if b < 0 {
		a *= -1
		b *= -1
	}
	if a > 0 {
		return (a-1)/b + 1
	}
	return a / b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func PowMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % M
		}
		a = a * a % M
		n /= 2
	}
	return res
}

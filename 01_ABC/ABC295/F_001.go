package main

import (
	"fmt"
	"strconv"
)

func f(s string, n int) int {
	res := 0
	si, _ := strconv.Atoi(s)
	for i := len(s) - 1; i < 16; i++ {
		p := pow(10, i+1)
		q := pow(10, i-len(s)+1)
		l := 0
		r := pow(10, 17-len(s))
		for l+1 < r {
			x := (l + r) / 2
			var v int
			if s[0] == '0' {
				v = x - 1 + q
			} else {
				v = x - 1
			}
			if (v/q)*p+si*q+v%q > n {
				r = x
			} else {
				l = x
			}
		}
		res += l
	}
	return res
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		var l, r int
		fmt.Scan(&s, &l, &r)
		fmt.Println(f(s, r) - f(s, l-1))
	}
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}

package main

import (
	"fmt"
	"math"
)

func ord(a, m int) int {
	a %= m
	if gcd(a, m) != 1 {
		return -1
	}
	sq := int(math.Sqrt(float64(m))) + 1
	mp := map[int]int{}
	s := a
	for i := 1; i <= sq; i++ {
		if _, ok := mp[s]; !ok {
			mp[s] = i
		}
		s = s * a % m
	}
	g := invMod(powMod(a, sq, m), m)
	x := 1
	for k := 0; k <= m/sq; k++ {
		if _, ok := mp[x]; ok {
			return sq*k + mp[x]
		}
		x = x * g % m
	}
	return 0
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		if m%2 == 0 {
			m /= 2
		}
		m *= 9
		fmt.Println(ord(10, m))
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func powMod(a, n, m int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % m
		}
		a = a * a % m
		n /= 2
	}
	return res
}

func invMod(a, m int) int {
	if a == 1 {
		return 1
	}
	return m + (1-m*invMod(m%a, a))/a
}

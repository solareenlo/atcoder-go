package main

import "fmt"

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)

	fmt.Println(base(a, k) * base(b, k))
}

func base(a, base int) int {
	res := 0
	cnt := 0
	for a > 0 {
		rem := a % 10
		res += rem * pow(base, cnt)
		a /= 10
		cnt++
	}
	return res
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

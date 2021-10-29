package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	l := 1
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
		a[i] /= 2
		l = lcm(l, a[i])
	}

	for i := 0; i < n; i++ {
		if l/a[i]%2 == 0 {
			fmt.Println(0)
			return
		}
	}

	res := m / (2 * l)
	if res >= l {
		res += 1
	}
	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

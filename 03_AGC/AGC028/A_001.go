package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	var s, t string
	fmt.Fscan(in, &n, &m, &s, &t)

	l := lcm(n, m)
	a := l / n
	b := l / m
	i := 0
	j := 0
	for i < n && j < m {
		if s[i] != t[j] {
			fmt.Println(-1)
			return
		}
		i += b
		j += a
	}
	fmt.Println(l)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}

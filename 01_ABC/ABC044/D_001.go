package main

import "fmt"

func f(b, n int) int {
	if n < b {
		return n
	}
	return f(b, n/b) + n%b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	res := int(1e13)
	if n == s {
		res = n + 1
	}

	for i := 1; i*i <= n-s; i++ {
		if f(i+1, n) == s {
			res = min(res, i+1)
		}
		if f((n-s)/i+1, n) == s {
			res = min(res, (n-s)/i+1)
		}
	}

	if res == int(1e13) {
		res = -1
	}
	fmt.Println(res)
}

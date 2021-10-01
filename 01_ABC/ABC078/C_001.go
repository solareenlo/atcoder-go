package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println((1900*m + 100*(n-m)) * pow(2, m))
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

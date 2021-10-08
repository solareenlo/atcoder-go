package main

import "fmt"

func main() {
	var d, n int
	fmt.Scan(&d, &n)

	x := pow(100, d)
	if n == 100 {
		fmt.Println(x * 101)
	} else {
		fmt.Println(x * n)
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

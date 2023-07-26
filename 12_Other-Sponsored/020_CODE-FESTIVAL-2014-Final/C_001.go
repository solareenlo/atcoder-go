package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 10; i < 10001; i++ {
		x := i
		memo := 0
		z := 0
		for x != 0 {
			memo += x % 10 * pow(i, z)
			z++
			x /= 10
		}
		if memo == N {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
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

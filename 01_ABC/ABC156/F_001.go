package main

import "fmt"

func main() {
	var k, q int
	fmt.Scan(&k, &q)

	d := make([]int, k)
	for i := range d {
		fmt.Scan(&d[i])
	}

	for i := 0; i < q; i++ {
		var n, x, m int
		fmt.Scan(&n, &x, &m)
		a := x % m
		for j := 0; j < k; j++ {
			a += ((d[j]+m-1)%m + 1) * ((n - 2 + k - j) / k)
		}
		fmt.Println(n - 1 - a/m)
	}
}

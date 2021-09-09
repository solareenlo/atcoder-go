package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	var l, r, t int
	res := make([]int, n)
	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r, &t)
		for j := l; j <= r; j++ {
			res[j-1] = t
		}
	}

	for i := range res {
		fmt.Println(res[i])
	}
}

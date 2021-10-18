package main

import "fmt"

func main() {
	var p, q, r int
	fmt.Scan(&p, &q, &r)

	fmt.Println(p + q + r - max(p, q, r))
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

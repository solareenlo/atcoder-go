package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		a[t]++
	}

	res := 1
	for i := n - 1; i > 0; i -= 2 {
		if a[i] != 2 {
			res = 0
		}
		res = res * 2 % int(1e9+7)
	}
	if n%2 == 1 && a[0] == 0 {
		res = 0
	}
	fmt.Println(res)
}

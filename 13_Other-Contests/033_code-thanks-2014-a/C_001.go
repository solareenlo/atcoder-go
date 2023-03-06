package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	ans := 0
	for m != 0 {
		var x int
		fmt.Scan(&x)
		ans += a[x]
		m--
	}
	fmt.Println(ans)
}

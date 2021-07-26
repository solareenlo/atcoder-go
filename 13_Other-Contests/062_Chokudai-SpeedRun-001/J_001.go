package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)

	a := make([]int, n+1)
	res := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		for j := m; j > 0; j -= (j & (-j)) {
			res -= a[j]
		}
		res += i
		for j := m; j <= n; j += (j & (-j)) {
			a[j]++
		}
	}
	fmt.Println(res)
}

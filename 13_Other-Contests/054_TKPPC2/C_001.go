package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		a[i] ^= 1
		a[i] += a[i-1]
	}

	l, r := 0, 1
	maxx := 0
	for l <= n && r <= n {
		l++
		for a[r]-a[l-1] <= m && r <= n {
			r++
		}
		if r-l > maxx {
			maxx = r - l
		}
	}

	fmt.Println(maxx)
}

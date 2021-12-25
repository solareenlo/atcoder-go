package main

import "fmt"

func main() {
	var n, h, w int
	fmt.Scan(&n, &h, &w)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if a[i] >= h && b[i] >= w {
			cnt++
		}
	}
	fmt.Println(cnt)
}

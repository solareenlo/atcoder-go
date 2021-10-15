package main

import "fmt"

func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)
	b := make([]int, m)
	for i := range b {
		fmt.Scan(&b[i])
	}

	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += b[j] * a[i][j]
		}
		if sum+c > 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}

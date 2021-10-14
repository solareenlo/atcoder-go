package main

import "fmt"

func main() {
	var n, m, k, tmp int
	fmt.Scan(&n, &m)
	a := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&k)
		for j := 0; j < k; j++ {
			fmt.Scan(&tmp)
			a[tmp-1]++
		}
	}

	cnt := 0
	for i := 0; i < m; i++ {
		if a[i] == n {
			cnt++
		}
	}
	fmt.Println(cnt)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	t := make([]int, n)
	for i := range t {
		fmt.Scan(&t[i])
	}
	var m int
	fmt.Scan(&m)
	p, x := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&p[i], &x[i])
	}

	for i := 0; i < m; i++ {
		res := 0
		for j := 0; j < n; j++ {
			if j == p[i]-1 {
				res += x[i]
			} else {
				res += t[j]
			}
		}
		fmt.Println(res)
	}
}

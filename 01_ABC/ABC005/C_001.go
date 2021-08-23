package main

import "fmt"

func main() {
	var t, n int
	fmt.Scan(&t, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	var m int
	fmt.Scan(&m)
	b := make([]int, m)
	for i := range b {
		fmt.Scan(&b[i])
	}

	res := "no"
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[j] <= b[i] && a[j]+t >= b[i] {
				cnt++
				a[j] = 200
				break
			}
		}
	}
	if cnt == m {
		res = "yes"
	}
	fmt.Println(res)
}

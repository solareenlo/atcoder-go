package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	cnt := 0
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if a[j]+p[i] == s {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

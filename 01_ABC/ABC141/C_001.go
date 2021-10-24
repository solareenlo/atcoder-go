package main

import "fmt"

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)

	p := make([]int, n)
	for i := 0; i < q; i++ {
		var a int
		fmt.Scan(&a)
		p[a-1]++
	}

	for i := 0; i < n; i++ {
		if k+p[i]-q > 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

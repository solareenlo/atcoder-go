package main

import "fmt"

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)

	sum := n * m
	for i := 0; i < n-1; i++ {
		var a int
		fmt.Scan(&a)
		sum -= a
	}

	if sum <= 0 {
		fmt.Println(0)
	} else if sum <= k {
		fmt.Println(sum)
	} else {
		fmt.Println(-1)
	}
}

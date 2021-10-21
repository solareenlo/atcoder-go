package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&a[i])
	}

	cnt := 0
	for i := n; i > 0; i-- {
		for j := i + i; j < n+1; j += i {
			a[i] ^= a[j]
		}
		if a[i] != 0 {
			cnt++
		}
	}
	fmt.Println(cnt)

	for i := 1; i < n+1; i++ {
		if a[i] != 0 {
			fmt.Println(i)
		}
	}
}

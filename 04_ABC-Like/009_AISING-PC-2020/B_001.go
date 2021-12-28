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
	for i := 1; i < n+1; i++ {
		if i%2 != 0 && a[i]%2 != 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}

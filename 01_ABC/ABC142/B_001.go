package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var h int
	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		if h >= k {
			cnt++
		}
	}

	fmt.Println(cnt)
}

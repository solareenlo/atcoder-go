package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	cnt, sum := 1, 0
	for i := 0; i < n; i++ {
		var l int
		fmt.Scan(&l)
		sum += l
		if sum <= x {
			cnt++
		}
	}
	fmt.Println(cnt)
}

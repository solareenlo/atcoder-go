package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	x *= 100

	cnt, sum := 0, 0
	for i := 0; i < n; i++ {
		var v, p int
		fmt.Scan(&v, &p)
		sum += v * p
		if sum <= x {
			cnt++
		}
	}

	if cnt == n {
		fmt.Println(-1)
	} else {
		fmt.Println(cnt + 1)
	}
}

package main

import "fmt"

func main() {
	var n, k, a int
	now := 0
	fmt.Scan(&n, &k)
	for i := 0; i < k; i++ {
		fmt.Scan(&a)
		now += a
		if now == n {
			break
		} else if now > n {
			now = 2 * n - now
		}
	}
	fmt.Println(now)
}

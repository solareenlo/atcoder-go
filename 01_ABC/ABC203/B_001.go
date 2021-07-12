package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			res += 100*i + j
		}
	}
	fmt.Println(res)
}

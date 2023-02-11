package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	res := 0
	for i := 0; i < n+1; i++ {
		tmp := sumOrder(i)
		if a <= tmp && tmp <= b {
			res += i
		}
	}
	fmt.Println(res)
}

func sumOrder(n int) int {
	if n == 0 {
		return 0
	}
	return sumOrder(n/10) + (n % 10)
}

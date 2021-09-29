package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res, l, r := 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r)
		res += r - l + 1
	}
	fmt.Println(res)
}

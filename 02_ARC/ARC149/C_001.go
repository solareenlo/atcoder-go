package main

import "fmt"

var n int
var ans []int

func R(x int) {
	for ; x <= n*n; x += 6 {
		ans = append(ans, x)
	}
}

func main() {
	fmt.Scan(&n)
	if n == 3 {
		fmt.Println(9, 3, 1)
		fmt.Println(5, 7, 8)
		fmt.Println(4, 2, 6)
		return
	}
	if n == 4 {
		fmt.Println(15, 11, 16, 12)
		fmt.Println(13, 3, 6, 9)
		fmt.Println(14, 7, 8, 1)
		fmt.Println(4, 2, 10, 5)
		return
	}
	R(2)
	R(4)
	R(6)
	R(3)
	R(1)
	R(5)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", ans[i*n+j])
		}
		fmt.Println()
	}
}

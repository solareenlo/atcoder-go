package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, 10000)
	for i := 1; i <= n; i++ {
		var k int
		fmt.Scan(&k)
		a[k] = 1
	}

	for i := 0; i >= 0; i++ {
		if a[i] == 0 {
			fmt.Println(i)
			return
		}
	}
}

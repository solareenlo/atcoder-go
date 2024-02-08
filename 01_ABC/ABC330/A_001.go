package main

import "fmt"

func main() {
	var n, l int
	fmt.Scan(&n, &l)
	num := 0
	var a [1010]int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] >= l {
			num++
		}
	}
	fmt.Println(num)
}

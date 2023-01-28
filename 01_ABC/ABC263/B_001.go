package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var p [51]int
	for i := 2; i <= n; i++ {
		var t int
		fmt.Scan(&t)
		p[i] = p[t] + 1
	}
	fmt.Println(p[n])
}

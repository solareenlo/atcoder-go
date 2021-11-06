package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n+1)
	res := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if i+a < n {
			p[i+a]++
		}
		if i-a >= 0 {
			res += p[i-a]
		}
	}

	fmt.Println(res)
}

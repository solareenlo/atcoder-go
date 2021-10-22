package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if i == 0 {
			res[0] += a
		}
		a = n - a
		for j := 1; j < n; j++ {
			res[n-j] = (res[n-j] + a) % n
			a = (a * i) % n
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(res[i], " ")
	}
}

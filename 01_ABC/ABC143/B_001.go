package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	d := make([]int, n)
	for i := range d {
		fmt.Scan(&d[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res += d[i] * d[j]
		}
	}

	fmt.Println(res)
}

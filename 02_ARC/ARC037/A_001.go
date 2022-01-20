package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := make([]int, n)
	for i := range m {
		fmt.Scan(&m[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		if m[i] < 80 {
			sum += 80 - m[i]
		}
	}

	fmt.Println(sum)
}

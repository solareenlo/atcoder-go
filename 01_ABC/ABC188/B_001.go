package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		res += a[i] * b[i]
	}

	if res != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

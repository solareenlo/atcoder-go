package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	pos := 0
	for i := 0; i < n; i++ {
		if abs(a[pos]) < abs(a[i]) {
			pos = i
		}
	}

	fmt.Println(2*n - 2)
	for i := 0; i < n; i++ {
		if i != pos {
			fmt.Println(pos+1, i+1)
		}
	}
	if a[pos] >= 0 {
		for i := 0; i < n-1; i++ {
			fmt.Println(i+1, i+2)
		}
	} else {
		for i := 0; i < n-1; i++ {
			fmt.Println(n-i, n-i-1)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package main

import "fmt"

func main() {
	var n int
	var va, vb, l float64
	fmt.Scan(&n, &va, &vb, &l)

	for i := 0; i < n; i++ {
		t := l / va
		l = vb * t
	}

	fmt.Println(l)
}

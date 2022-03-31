package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	t := 0
	for i := 0; i < n; i++ {
		var h int
		fmt.Scan(&h)
		if t < h {
			t = h
		} else {
			break
		}
	}
	fmt.Println(t)
}

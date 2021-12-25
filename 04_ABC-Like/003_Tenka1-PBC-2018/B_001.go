package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	for i := 0; i < k; i++ {
		if i%2 == 0 {
			b += a / 2
			a /= 2
		} else {
			a += b / 2
			b /= 2
		}
	}

	fmt.Println(a, b)
}

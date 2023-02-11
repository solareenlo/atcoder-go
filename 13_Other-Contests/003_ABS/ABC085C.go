package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	y /= 1000

	a, b, c := -1, -1, -1
	for i := 0; i < n+1; i++ {
		for j := 0; j < n-i+1; j++ {
			if i*10+j*5+n-i-j == y {
				a = i
				b = j
				c = n - i - j
				break
			}
		}
	}
	fmt.Printf("%d %d %d\n", a, b, c)
}

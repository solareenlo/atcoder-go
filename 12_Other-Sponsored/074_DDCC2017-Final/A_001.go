package main

import "fmt"

var k int

func main() {
	fmt.Scan(&k)
	fmt.Println(w(200), w(300))
}

func w(n int) int {
	r := n / 2
	ans := 0
	var x, y int
	for i := -r; i < r; i += k {
		for j := -r; j <= r; j += k {
			if i < 0 {
				x = i
			} else {
				x = i + k
			}
			if j < 0 {
				y = j
			} else {
				y = j + k
			}
			if x*x+y*y <= r*r {
				ans++
			}
		}
	}
	return ans
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for k := 0; k < 30; k++ {
		x := 1
		for i := 0; i < 30; i++ {
			x = x * 3
			if i == k {
				x++
			}
		}
		if n == x {
			fmt.Println(k + 1)
			return
		}
	}
	fmt.Println(-1)
}

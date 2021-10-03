package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	y /= 1000

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if n-i-j >= 0 {
				if i*10+j*5+(n-i-j) == y {
					fmt.Println(i, j, n-i-j)
					return
				}
			}
		}
	}
	fmt.Println(-1, -1, -1)
}

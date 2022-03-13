package main

import "fmt"

func main() {
	var m, d int
	fmt.Scan(&m, &d)

	ans := 0
	for i := 1; i <= d; i++ {
		x := i % 10
		y := i / 10
		if x >= 2 && y >= 2 && x*y <= m {
			ans++
		}
	}
	fmt.Println(ans)
}

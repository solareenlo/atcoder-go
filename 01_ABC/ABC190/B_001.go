package main

import "fmt"

func main() {
	var n, s, d int
	fmt.Scan(&n, &s, &d)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if x < s && d < y {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

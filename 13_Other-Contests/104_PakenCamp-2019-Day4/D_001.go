package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	c := 0
	for i := 0; i < 1010; i++ {
		k := (i + 1) * (i + 2) / 2
		if k <= x {
			c = i + 1
		}
	}
	y := x - c*(c+1)/2
	fmt.Println(2000)
	for i := 0; i < c; i++ {
		fmt.Print("1 ")
	}
	if c != n {
		fmt.Print(2001-y, " ")
	}
	for i := c + 1; i < n; i++ {
		fmt.Print("2001 ")
	}
	fmt.Println()
}

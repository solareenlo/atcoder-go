package main

import "fmt"

func main() {
	var a, b, c, d, r int
	fmt.Scan(&a, &b, &c, &d, &r)
	x := (b + d - 1) / d * d
	x = max(x, c)
	if c+r <= x {
		fmt.Println("No")
	} else if c == x {
		fmt.Println("Yes")
	} else if a < c && x <= a+r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

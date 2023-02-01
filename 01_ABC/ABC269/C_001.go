package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for r := x; ; r = (r - 1) & x {
		fmt.Println(x - r)
		if r == 0 {
			break
		}
	}
}

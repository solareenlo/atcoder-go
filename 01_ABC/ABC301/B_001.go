package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a int
	fmt.Scan(&a)
	fmt.Println(a)
	for n-1 > 0 {
		n--
		var b int
		fmt.Scan(&b)
		for a != b {
			if a < b {
				a++
			} else {
				a--
			}
			fmt.Printf(" %d", a)
		}
	}
	fmt.Println()
}

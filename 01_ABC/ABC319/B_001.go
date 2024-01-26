package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= 9; j++ {
			if n%j == 0 && i%(n/j) == 0 {
				fmt.Print(j)
				break
			}
			if j == 9 {
				fmt.Print("-")
			}
		}
	}
}

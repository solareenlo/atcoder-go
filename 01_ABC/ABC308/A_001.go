package main

import "fmt"

func main() {
	last := 0
	for i := 0; i < 8; i++ {
		var x int
		fmt.Scan(&x)
		if x < last || x < 100 || x > 675 || x%25 != 0 {
			fmt.Println("No")
			return
		}
		last = x
	}
	fmt.Println("Yes")
}

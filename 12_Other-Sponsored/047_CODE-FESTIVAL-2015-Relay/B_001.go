package main

import (
	"fmt"
)

func main() {
	var a [20]int
	for i := 1; i <= 10; i++ {
		var b string
		fmt.Scan(&b)
		for j := 1; j <= 10; j++ {
			if b[j-1] == 'o' {
				a[j] = 1
			}
		}
	}
	for i := 1; i <= 10; i++ {
		if a[i] == 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

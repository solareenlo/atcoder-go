package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j == n {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}

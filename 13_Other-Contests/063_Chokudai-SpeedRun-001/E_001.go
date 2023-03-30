package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a == 1 {
			fmt.Println(i + 1)
			return
		}
	}
}

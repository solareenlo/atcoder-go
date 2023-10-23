package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := a; i >= 1; i-- {
		if i%b == 0 {
			fmt.Println(i)
			return
		}
	}
}

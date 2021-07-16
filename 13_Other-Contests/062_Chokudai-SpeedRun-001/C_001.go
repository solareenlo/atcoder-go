package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a int
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a)
		fmt.Print(a, ",")
	}
	fmt.Scan(&a)
	fmt.Println(a)
}

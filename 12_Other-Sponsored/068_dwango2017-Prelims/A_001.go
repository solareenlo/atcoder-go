package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	if a+b <= n {
		fmt.Println(0)
		return
	} else {
		fmt.Println(a + b - n)
	}
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%100 == 0 {
		fmt.Println(n / 100)
	} else {
		fmt.Println(n/100 + 1)
	}
}

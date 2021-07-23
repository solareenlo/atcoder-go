package main

import "fmt"

func main() {
	var r int
	fmt.Scan(&r)
	switch {
	case r < 1200:
		fmt.Println("ABC")
	case r < 2800:
		fmt.Println("ARC")
	default:
		fmt.Println("AGC")
	}
}

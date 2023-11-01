package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	z := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < n; i++ {
		fmt.Print(string(z[i]))
	}
}

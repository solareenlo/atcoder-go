package main

import "fmt"

func main() {
	var a [300]int
	for i := 1; i <= 28; i++ {
		var n int
		fmt.Scan(&n)
		a[n]++
	}
	for i := 1; i <= 30; i++ {
		if a[i] == 0 {
			fmt.Println(i)
		}
	}
}

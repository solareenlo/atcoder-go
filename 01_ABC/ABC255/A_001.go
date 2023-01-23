package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	var a [4]int
	for i, _ := range a {
		fmt.Scan(&a[i])
	}

	fmt.Println(a[2*r+c-3])
}

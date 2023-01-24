package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(1 << x)
}

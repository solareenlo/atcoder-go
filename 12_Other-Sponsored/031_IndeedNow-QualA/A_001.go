package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(len(a) * len(b))
}

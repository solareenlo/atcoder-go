package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println("0" + s[:3])
}

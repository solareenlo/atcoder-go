package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res := 0
	if a%b != 0 {
		res = b - a%b
	}
	fmt.Println(res)
}

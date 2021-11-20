package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	res := 4
	if a+b >= 15 && b >= 8 {
		res = 1
	} else if a+b >= 10 && b >= 3 {
		res = 2
	} else if a+b >= 3 {
		res = 3
	}

	fmt.Println(res)
}

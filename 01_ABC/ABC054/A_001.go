package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	res := "Alice"
	switch {
	case a == b:
		res = "Draw"
	case a == 1:
		res = "Alice"
	case b == 1:
		res = "Bob"
	case a < b:
		res = "Bob"
	}
	fmt.Println(res)
}

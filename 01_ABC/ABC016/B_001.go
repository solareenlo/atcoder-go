package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	res := "-"
	if a+b == c && a-b == c {
		res = "?"
	} else if a+b != c && a-b != c {
		res = "!"
	} else if a+b == c {
		res = "+"
	}
	fmt.Println(res)
}

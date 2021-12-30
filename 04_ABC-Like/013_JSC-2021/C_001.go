package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	c := b
	for ; ; c-- {
		if (a-1+c)/c < b/c {
			break
		}
	}

	fmt.Println(c)
}

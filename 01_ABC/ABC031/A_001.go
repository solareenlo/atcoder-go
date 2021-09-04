package main

import "fmt"

func main() {
	var a, d int
	fmt.Scan(&a, &d)

	if a > d {
		fmt.Println(a * (d + 1))
	} else {
		fmt.Println((a + 1) * d)
	}
}

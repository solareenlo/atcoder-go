package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	for i := 1; i <= 127; i++ {
		if (i%3 == a) && (i%5 == b) && (i%7 == c) {
			fmt.Println(i)
		}
	}
}

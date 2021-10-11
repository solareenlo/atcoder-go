package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	if k%2 != 0 {
		fmt.Println((k / 2) * (k/2 + 1))
	} else {
		fmt.Println((k / 2) * (k / 2))
	}
}

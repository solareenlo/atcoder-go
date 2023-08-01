package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n--
	if n%40 < 20 {
		fmt.Println(n%40 + 1)
	} else {
		fmt.Println(40 - n%40)
	}
}

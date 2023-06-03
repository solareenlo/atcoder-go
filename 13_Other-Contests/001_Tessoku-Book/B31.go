package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n/3 + n/5 + n/7 - n/15 - n/35 - n/21 + n/105)
}

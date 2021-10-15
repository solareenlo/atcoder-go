package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	t := "2019/04/30"
	if s <= t {
		fmt.Println("Heisei")
	} else {
		fmt.Println("TBD")
	}
}

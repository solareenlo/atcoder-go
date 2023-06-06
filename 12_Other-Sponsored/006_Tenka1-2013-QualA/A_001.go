package main

import "fmt"

func main() {
	s := 42
	for s <= 130000000 {
		s = s * 2
	}
	fmt.Println(s)
}

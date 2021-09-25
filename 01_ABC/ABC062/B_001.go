package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([]string, h)
	for i := range s {
		fmt.Scan(&s[i])
	}

	for i := 0; i < w+2; i++ {
		fmt.Print("#")
	}
	fmt.Println()
	for i := 0; i < h; i++ {
		fmt.Print("#", s[i], "#", "\n")
	}
	for i := 0; i < w+2; i++ {
		fmt.Print("#")
	}
	fmt.Println()
}

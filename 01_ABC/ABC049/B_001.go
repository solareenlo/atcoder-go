package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
	}

	for i := 0; i < h; i++ {
		fmt.Println(s[i])
		fmt.Println(s[i])
	}
}

package main

import "fmt"

func main() {
	var H, W, h, w int
	fmt.Scan(&H, &W, &h, &w)

	if h >= H && w <= W {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

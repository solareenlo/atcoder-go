package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	fmt.Println(h*(w-1) + (h-1)*w)
}

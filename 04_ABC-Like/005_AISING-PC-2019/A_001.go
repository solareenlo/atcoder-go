package main

import "fmt"

func main() {
	var n, h, w int
	fmt.Scan(&n, &h, &w)

	fmt.Println((n - h + 1) * (n - w + 1))
}

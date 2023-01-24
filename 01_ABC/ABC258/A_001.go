package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	fmt.Printf("%d:%02d", 21+t/60, t%60)
}

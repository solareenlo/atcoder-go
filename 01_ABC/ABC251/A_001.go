package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < 6/len(s); i++ {
		fmt.Print(s)
	}
}

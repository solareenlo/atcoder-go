package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		fmt.Println(i + 1)
	}
}

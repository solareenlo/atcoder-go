package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	if s == "hi" || s == "hihi" || s == "hihihi" || s == "hihihihi" || s == "hihihihihi" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

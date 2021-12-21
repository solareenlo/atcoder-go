package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	a := "oxxoxxoxxoxx"
	if strings.Index(a, s) == -1 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

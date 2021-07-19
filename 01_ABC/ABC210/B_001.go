package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	if strings.Index(s, "1")%2 == 1 {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Takahashi")
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	res := 0
	var s string
	for i := 0; i < 12; i++ {
		fmt.Scan(&s)
		if strings.Contains(s, "r") {
			res++
		}
	}
	fmt.Println(res)
}

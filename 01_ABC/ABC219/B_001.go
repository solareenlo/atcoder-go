package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := make([]string, 3)
	for i := range s {
		fmt.Scan(&s[i])
	}

	var t string
	fmt.Scan(&t)
	n := len(t)
	for i := 0; i < n; i++ {
		j, _ := strconv.Atoi(string(t[i]))
		fmt.Print(s[j-1])
	}
	fmt.Println()
}

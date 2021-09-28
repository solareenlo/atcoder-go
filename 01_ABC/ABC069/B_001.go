package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Printf(string(s[0]) + strconv.Itoa(len(s)-2) + string(s[len(s)-1]))
}

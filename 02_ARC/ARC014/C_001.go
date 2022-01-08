package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	res := 0
	res += strings.Count(s, "R") % 2
	res += strings.Count(s, "G") % 2
	res += strings.Count(s, "B") % 2

	fmt.Println(res)
}

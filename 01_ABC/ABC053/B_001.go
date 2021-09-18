package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	a, z := 0, 0
	a = strings.Index(s, "A")
	z = strings.LastIndex(s, "Z")
	fmt.Println(z - a + 1)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(strings.Replace(s, "HAGIYA", "HAGIXILE", 1))
}

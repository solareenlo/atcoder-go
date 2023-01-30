package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := "?rheuo"
	fmt.Println(strings.IndexByte(t, s[1]))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var a, b int
	fmt.Scan(&s, &a, &b)
	t := strings.Split(s, "")

	t[a-1], t[b-1] = t[b-1], t[a-1]
	fmt.Println(strings.Join(t, ""))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var k int
	fmt.Scan(&k)
	fmt.Println(strings.Repeat("ACL", k))
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scan(&s)

	reg := regexp.MustCompile(`^(dream|dreamer|erase|eraser)*$`)
	if reg.MatchString(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

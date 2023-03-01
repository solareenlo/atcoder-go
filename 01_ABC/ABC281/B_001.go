package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := regexp.MustCompile(`^([A-Z][1-9][0-9]{5}[A-Z])`)
	if r.MatchString(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

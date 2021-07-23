package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := regexp.MustCompile(`^A[a-z][a-z]*C[a-z]*[a-z]$`)
	if r.MatchString(s) {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}

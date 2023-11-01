package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scan(&s)
	re := regexp.MustCompile("00")
	s = re.ReplaceAllString(s, "0")
	fmt.Println(len(s))
}

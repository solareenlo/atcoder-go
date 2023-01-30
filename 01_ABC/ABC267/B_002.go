package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := []byte{s[0], s[6], s[3], max(s[1], s[7]), s[4], max(s[8], s[2]), s[5], s[9]}

	r := regexp.MustCompile(`^0.*10+1.*`)
	if r.MatchString(string(t[:])) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

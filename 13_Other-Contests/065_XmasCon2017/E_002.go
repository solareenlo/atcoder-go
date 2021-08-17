package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	s = strings.Replace(s, "", "B*", -1)
	s = strings.Replace(s, "A", "A?", -1)

	if regexp.MustCompile(`^` + s + `$`).MatchString(t) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

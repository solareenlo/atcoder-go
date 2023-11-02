package main

import (
	"fmt"
	"regexp"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	re := regexp.MustCompile("na")
	replaced := re.ReplaceAllString(S, "nya")

	fmt.Println(replaced)
}

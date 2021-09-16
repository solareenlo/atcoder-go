package main

import (
	"fmt"
	"strings"
)

func main() {
	var c string
	fmt.Scan(&c)
	if strings.Contains("aiueo", c) {
		fmt.Println("vowel")
	} else {
		fmt.Println("consonant")
	}
}

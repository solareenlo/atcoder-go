package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	f := []string{"Oo", "o", "O"}
	for i := 100; i > 2; i-- {
		pos := strings.Index(s, strings.Repeat("o", i)+"kayama")
		for pos != -1 {
			s = s[:pos] + f[i%3] + s[pos+i:]
			pos = strings.Index(s, strings.Repeat("o", i)+"kayama")
		}
	}
	fmt.Println(s)
}

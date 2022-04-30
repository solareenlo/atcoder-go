package main

import (
	"fmt"
	"unicode"
)

func main() {
	var S string
	var X int
	fmt.Scan(&S, &X)
	X--
	for {
		sz := 0
		for _, c := range S {
			if unicode.IsDigit(c) {
				t := int(c - '0')
				if sz*(t+1) > X {
					X %= sz
					break
				}
				sz *= t + 1
			} else {
				if sz == X {
					fmt.Println(string(c))
					return
				}
				sz++
			}
		}
	}
}

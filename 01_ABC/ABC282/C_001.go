package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	var tmp string
	fmt.Scan(&tmp)
	S := strings.Split(tmp, "")
	flag := false
	for i := 0; i < N; i++ {
		if S[i] == "\"" {
			if flag {
				flag = false
			} else {
				flag = true
			}
		} else if S[i] == "," {
			if !flag {
				S[i] = "."
			}
		}
	}
	fmt.Println(strings.Join(S, ""))
}

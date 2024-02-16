package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	s := strings.Split(S, "")
	for i := 0; i < len(s); i++ {
		s[i] = string((s[i][0]-'A'+23)%26 + 'A')
	}
	fmt.Println(strings.Join(s, ""))
}

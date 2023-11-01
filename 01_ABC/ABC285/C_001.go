package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	x := 0
	for i := 0; i < len(s); i++ {
		x = x*26 + int(s[i]-'A'+1)
	}
	fmt.Println(x)
}

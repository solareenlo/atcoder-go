package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 45
	for i := range s {
		ans -= int(s[i] - '0')
	}
	fmt.Println(ans)
}

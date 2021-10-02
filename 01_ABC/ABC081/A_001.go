package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	res := 0
	for i := 0; i < 3; i++ {
		res += int(s[i] - '0')
	}
	fmt.Println(res)
}

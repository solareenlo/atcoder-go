package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, 10)
	s[0] = "pon"
	s[1] = "pon"
	s[2] = "hon"
	s[3] = "bon"
	s[4] = "hon"
	s[5] = "hon"
	s[6] = "pon"
	s[7] = "hon"
	s[8] = "pon"
	s[9] = "hon"

	fmt.Println(s[n%10])
}

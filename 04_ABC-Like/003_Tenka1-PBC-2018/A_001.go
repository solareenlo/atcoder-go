package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	if n == 3 {
		s = reverseString(s)
	}
	fmt.Println(s)
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

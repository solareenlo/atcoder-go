package main

import "fmt"

var s string

func main() {
	var n int
	fmt.Scan(&n)
	n--

	rec(n)
	fmt.Println(reverseString(s))
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func rec(n int) {
	if n/26 == 0 {
		c := 'a' + (n % 26)
		s += string(c)
		return
	}
	rec(n % 26)
	rec(n/26 - 1)
}

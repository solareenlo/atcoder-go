package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	t := reverseString(s)
	if s == t {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

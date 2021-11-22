package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	if len(n) == 1 {
		fmt.Println("Yes")
		return
	}

	for n[len(n)-1] == '0' {
		n = n[:len(n)-1]
	}
	m := n
	m = reverseString(m)

	if n == m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

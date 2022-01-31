package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			fmt.Println(i, i+1)
			return
		}
	}

	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			fmt.Println(i-1, i+1)
			return
		}
	}

	fmt.Println(-1, -1)
}

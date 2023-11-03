package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	if n == 1 {
		fmt.Println("Yes")
		return
	}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

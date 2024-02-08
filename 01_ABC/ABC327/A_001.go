package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	for i := 0; i < len(s)-1; i++ {
		if (s[i] == 'a' && s[i+1] == 'b') || (s[i] == 'b' && s[i+1] == 'a') {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

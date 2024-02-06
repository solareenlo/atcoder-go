package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	for i := 0; i < n-2; i++ {
		if s[i] == 'A' && s[i+1] == 'B' && s[i+2] == 'C' {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}

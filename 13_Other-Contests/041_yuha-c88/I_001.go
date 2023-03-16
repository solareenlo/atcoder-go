package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == 'm' || s[i][j] == 'o' || s[i][j] == 'n' || s[i][j] == 'i' || s[i][j] == 'c' || s[i][j] == 'a' {
				fmt.Printf("%c", s[i+1][j])
			}
		}
	}
	fmt.Println()
}

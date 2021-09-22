package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	b := make([]int, 26)
	for i := range b {
		b[i] = 50
	}

	var s string
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		a := make([]int, 26)
		for j := range s {
			a[s[j]-'a']++
		}
		for j := 0; j < 26; j++ {
			b[j] = min(b[j], a[j])
		}
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < b[i]; j++ {
			fmt.Printf("%c", i+'a')
		}
	}
	fmt.Println()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

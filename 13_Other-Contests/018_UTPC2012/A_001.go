package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	a := make([]int, 10)
	for i := 0; i < 4; i++ {
		a[s[i]-'0']++
	}
	b := make([]int, 10)
	for i := 5; i < len(s); i++ {
		if s[i] != '/' {
			b[s[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		if a[i] != b[i] {
			fmt.Println("no")
			return
		}
	}
	fmt.Println("yes")
}

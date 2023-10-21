package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a, b, c, d := 0, 0, 0, 0
	for i := 0; i < 5; i++ {
		if s[i] == 'a' {
			a++
		}
		if s[i] == 'y' {
			b++
		}
		if s[i] == 'h' {
			c++
		}
		if s[i] == 'o' {
			d++
		}
	}
	if b == 1 && a == 1 && c == 1 && d == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

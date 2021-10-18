package main

import "fmt"

func main() {
	var l string
	fmt.Scan(&l)

	mod := int(1e9 + 7)
	a, b := 0, 1
	for _, c := range l {
		a = (a * 3) % mod
		if c == '1' {
			a = (a + b) % mod
			b = (b * 2) % mod
		}
	}
	fmt.Println((a + b) % mod)
}

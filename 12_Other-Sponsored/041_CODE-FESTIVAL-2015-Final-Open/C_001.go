package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	s += "#"
	tmp := 0
	for i := 0; i < len(s); i++ {
		if (s[i] == '0' && s[i+1] == '1') || (s[i] == '1' && s[i+1] == '0') {
			i++
			tmp++
		}
	}
	fmt.Println(n - tmp)
}

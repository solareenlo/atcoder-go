package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			a = append(a, i)
		} else {
			fmt.Println(a[len(a)-1]+1, i+1)
			a = a[:len(a)-1]
		}
	}
}

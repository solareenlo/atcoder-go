package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := make([]int, 0)
	r := 0
	var x [100100]int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			a = append(a, i)
		} else if len(a) != 0 {
			x[i+1] = x[a[len(a)-1]] + 1
			r += x[i+1]
			a = a[:len(a)-1]
		}
	}
	fmt.Println(r)
}

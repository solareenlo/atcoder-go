package main

import "fmt"

func main() {
	var l int
	fmt.Scan(&l)
	var s, t string
	fmt.Scan(&s, &t)
	if s+t > t+s {
		s, t = t, s
	}
	tmp := ""
	for l%len(s) != 0 {
		tmp += t
		l -= len(t)
	}
	for i := 0; i < (l / len(s)); i++ {
		fmt.Print(s)
	}
	fmt.Println(tmp)
}

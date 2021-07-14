package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	res := 0
	for i := 0; i < 10000; i++ {
		a := i % 10
		b := i / 10 % 10
		c := i / 100 % 10
		d := i / 1000 % 10
		flag := true
		for j := range s {
			e := (a == j) || (b == j) || (c == j) || (d == j)
			if (e && s[j] == 'x') || (!e && s[j] == 'o') {
				flag = false
			}
		}
		if flag {
			res++
		}
	}
	fmt.Println(res)
}

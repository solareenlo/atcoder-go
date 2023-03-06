package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	s := make([]string, n)
	for i := range s {
		s[i] = "0"
	}
	if d%3 != 0 {
		for i := 0; i < n; i++ {
			if i%3 == 0 {
				s[i] = "1"
			}
		}
		fmt.Println(strings.Join(s, ""))
		a := -1
		b := -1
		for i := 0; i < 3; i++ {
			var x, y string
			fmt.Scan(&x, &y)
			if x == "1" {
				a = i
			}
			if y == "1" {
				b = i
			}
			if i < 2 {
				fmt.Println("Move(A,-1)")
				fmt.Scan(&x, &y)
				fmt.Println("Move(B,-1)")
			}
		}
		if (a+d)%3 == b%3 {
			fmt.Println("i<j")
		} else {
			fmt.Println("i>j")
		}
		return
	}
	t := 1
	for d%3 == 0 {
		d /= 3
		t *= 3
	}
	for i := 0; i < n; {
		for j := 0; j < t; j++ {
			if i < n {
				s[i] = "0"
				i++
			}
		}
		for j := 0; j < t; j++ {
			if i < n {
				s[i] = "1"
				i++
			}
		}
		for j := 0; j < t; j++ {
			if i < n {
				s[i] = string(int("1"[0]) - j%2)
				i++
			}
		}
	}
	fmt.Println(strings.Join(s, ""))

	var a, b string
	for i := 0; i < 3; i++ {
		var x, y string
		fmt.Scan(&x, &y)
		a += x
		b += y
		if i < 2 {
			fmt.Println("Move(A,1)")
			fmt.Scan(&x, &y)
			fmt.Println("Move(B,1)")
		}
	}
	ia := get(a)
	ib := get(b)
	if (ia+d)%3 == ib%3 {
		fmt.Println("i<j")
	} else {
		fmt.Println("i>j")
	}
}

func get(s string) int {
	if s == "000" || s == "001" || s == "011" {
		return 0
	}
	if s == "111" || s == "110" {
		return 1
	}
	return 2
}

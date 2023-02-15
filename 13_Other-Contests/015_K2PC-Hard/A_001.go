package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	ans := 1
	for n-ans > 0 {
		n -= ans
		ans++
	}
	ans++
	b := n
	a := ans - n
	ans = 1
	for m-ans > 0 {
		m -= ans
		ans++
	}
	ans++
	b1 := m
	a1 := ans - m
	x := a + a1
	y := b + b1
	ans = 1
	k := 0
	for ans+1 < x+y {
		k += ans
		ans++
	}
	k += y
	fmt.Println(k)
}

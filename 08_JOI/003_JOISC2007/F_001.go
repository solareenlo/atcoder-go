package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var a [26]int
	for i := 0; i < len(s); i++ {
		a[s[i]-'A']++
	}

	res := 0
	for i := 0; i < len(s); i++ {
		c := 0
		x := int(s[i] - 'A')
		for j := 0; j < x; j++ {
			c += a[j]
		}
		c *= f(len(s) - i - 1)
		for j := 0; j < 26; j++ {
			c /= f(a[j])
		}
		res += c
		a[x]--
	}
	fmt.Println(res + 1)
}

func f(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

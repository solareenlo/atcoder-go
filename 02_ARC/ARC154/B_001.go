package main

import "fmt"

func main() {
	var d [256]int

	var n int
	var a, b string
	fmt.Scan(&n, &a, &b)
	k := n - 1
	i := k
	j := k
	for k >= 0 {
		d[a[k]]++
		d[b[k]]--
		k--
	}
	c := 0
	for t := 'a'; t <= 'k'; t++ {
		if d[t] != 0 {
			c = -1
		}
	}
	if c != -1 {
		for j >= 0 {
			if a[i] == b[j] {
				i--
				j--
			} else {
				j--
				c++
			}
		}
	}
	fmt.Println(c)
}

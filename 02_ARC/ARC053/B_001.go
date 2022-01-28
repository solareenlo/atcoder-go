package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	alpha := map[byte]int{}
	for i := 0; i < n; i++ {
		alpha[s[i]]++
	}

	odd := 0
	for _, v := range alpha {
		if v%2 != 0 {
			odd++
		}
	}

	if odd == 0 {
		fmt.Println(n)
	} else {
		fmt.Println(2*((n-odd)/(2*odd)) + 1)
	}
}

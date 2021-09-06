package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	p := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i], &p[i])
		sum += p[i]
	}

	res := -1
	for i := 0; i < n; i++ {
		if p[i] > sum/2 {
			res = i
			break
		}
	}
	if res == -1 {
		fmt.Println("atcoder")
	} else {
		fmt.Println(s[res])
	}
}
